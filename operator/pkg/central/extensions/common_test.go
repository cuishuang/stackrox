package extensions

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	platform "github.com/stackrox/rox/operator/apis/platform/v1alpha1"
	"github.com/stackrox/rox/operator/pkg/utils/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	ctrlClient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
)

type secretVerifyFunc func(t *testing.T, central *platform.Central, secret *v1.Secret)
type statusVerifyFunc func(t *testing.T, status *platform.CentralStatus)

type secretReconciliationTestCase struct {
	Spec                   platform.CentralSpec
	Deleted                bool
	Existing               []*v1.Secret
	ExistingManaged        []*v1.Secret
	Other                  []ctrlClient.Object
	InterceptedK8sAPICalls interceptor.Funcs

	ExpectedCreatedSecrets       map[string]secretVerifyFunc
	ExpectedSecretsAfterDeletion []string
	ExpectedError                string
	ExpectedNotExistingSecrets   []string
	VerifyStatus                 statusVerifyFunc
}

func basicSpecWithScanner(scannerEnabled bool, scannerV4Enabled bool) platform.CentralSpec {
	spec := platform.CentralSpec{
		Scanner: &platform.ScannerComponentSpec{
			ScannerComponent: new(platform.ScannerComponentPolicy),
		},
		ScannerV4: &platform.ScannerV4Spec{
			ScannerComponent: new(platform.ScannerV4ComponentPolicy),
		},
	}
	if scannerEnabled {
		*spec.Scanner.ScannerComponent = platform.ScannerComponentEnabled
	} else {
		*spec.Scanner.ScannerComponent = platform.ScannerComponentDisabled
	}

	if scannerV4Enabled {
		*spec.ScannerV4.ScannerComponent = platform.ScannerV4ComponentEnabled
	} else {
		*spec.ScannerV4.ScannerComponent = platform.ScannerV4ComponentDisabled
	}

	return spec
}

type runFn func(ctx context.Context, central *platform.Central, client ctrlClient.Client, direct ctrlClient.Reader, statusUpdater func(updateStatusFunc), log logr.Logger) error

func testSecretReconciliation(t *testing.T, runFn runFn, c secretReconciliationTestCase) {
	central := &platform.Central{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "platform.stackrox.io/v1alpha1",
			Kind:       "Central",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-central",
			Namespace: testutils.TestNamespace,
		},
		Spec: *c.Spec.DeepCopy(),
	}

	if c.Deleted {
		central.DeletionTimestamp = new(metav1.Time)
		*central.DeletionTimestamp = metav1.Now()
	}

	statusUpdater := func(statusFunc updateStatusFunc) {
		statusFunc(&central.Status)
	}

	var existingSecrets []ctrlClient.Object
	for _, existingSecret := range c.Existing {
		existingSecrets = append(existingSecrets, existingSecret.DeepCopy())
	}
	for _, existingManagedSecret := range c.ExistingManaged {
		managedSecret := existingManagedSecret.DeepCopy()
		managedSecret.SetOwnerReferences([]metav1.OwnerReference{*metav1.NewControllerRef(central, central.GroupVersionKind())})
		existingSecrets = append(existingSecrets, managedSecret)
	}
	var otherExisting []runtime.Object
	for _, existingObj := range c.Other {
		otherExisting = append(otherExisting, existingObj.DeepCopyObject())
	}

	sch := runtime.NewScheme()
	require.NoError(t, platform.AddToScheme(sch))
	require.NoError(t, scheme.AddToScheme(sch))
	client := fake.NewClientBuilder().
		WithScheme(sch).
		WithObjects(existingSecrets...).
		WithRuntimeObjects(otherExisting...).
		Build()

	client = interceptor.NewClient(client, c.InterceptedK8sAPICalls)

	// Verify that an initial invocation does not touch any of the existing unmanaged secrets, and creates
	// the expected managed ones.
	err := runFn(context.Background(), central.DeepCopy(), client, client, statusUpdater, logr.Discard())
	if c.ExpectedError == "" {
		require.NoError(t, err)
	} else {
		require.Error(t, err)
		assert.Contains(t, err.Error(), c.ExpectedError)
		return
	}

	if c.VerifyStatus != nil {
		c.VerifyStatus(t, &central.Status)
	}

	secretsList := &v1.SecretList{}
	err = client.List(context.Background(), secretsList, ctrlClient.InNamespace(testutils.TestNamespace))
	require.NoError(t, err)

	secretsByName := make(map[string]v1.Secret)
	for _, secret := range secretsList.Items {
		secretsByName[secret.Name] = secret
	}

	for _, existingSecret := range c.Existing {
		found, ok := secretsByName[existingSecret.Name]
		if !assert.Truef(t, ok, "pre-existing unmanaged secret %s is gone", existingSecret.Name) {
			continue
		}
		assert.Equalf(t, existingSecret.Data, found.Data, "data of pre-existing unmanaged secret %s has changed", existingSecret.Name)
		delete(secretsByName, existingSecret.Name)
	}

	for _, notExistingSecret := range c.ExpectedNotExistingSecrets {
		_, ok := secretsByName[notExistingSecret]
		assert.Falsef(t, ok, "secret %s was created", notExistingSecret)
	}

	for name, verifyFunc := range c.ExpectedCreatedSecrets {
		found, ok := secretsByName[name]
		if !assert.True(t, ok, "expected secret %s was not created", name) {
			continue
		}
		verifyFunc(t, central, &found)
		delete(secretsByName, name)
	}

	assert.Empty(t, secretsByName, "one or more unexpected secrets exist")

	// Verify that a second invocation does not further change the cluster state
	err = runFn(context.Background(), central.DeepCopy(), client, client, statusUpdater, logr.Discard())
	assert.NoError(t, err, "second invocation of reconciliation function failed")

	if c.VerifyStatus != nil {
		c.VerifyStatus(t, &central.Status)
	}

	secretsList2 := &v1.SecretList{}
	err = client.List(context.Background(), secretsList2, ctrlClient.InNamespace(testutils.TestNamespace))
	require.NoError(t, err)

	assert.ElementsMatch(t, secretsList.Items, secretsList2.Items, "second invocation changed the cluster state")

	// Fake deletion of the CR
	central.DeletionTimestamp = new(metav1.Time)
	*central.DeletionTimestamp = metav1.Now()

	err = runFn(context.Background(), central.DeepCopy(), client, client, statusUpdater, logr.Discard())
	assert.NoError(t, err, "deletion of CR resulted in error")

	secretsList3 := &v1.SecretList{}
	err = client.List(context.Background(), secretsList3, ctrlClient.InNamespace(testutils.TestNamespace))
	require.NoError(t, err)

	postDeletionSecretsByName := make(map[string]v1.Secret)
	for _, secret := range secretsList3.Items {
		postDeletionSecretsByName[secret.Name] = secret
	}

	// Verify pre-existing secrets still exist
	for _, existingSecret := range c.Existing {
		found, ok := postDeletionSecretsByName[existingSecret.Name]
		if !assert.Truef(t, ok, "pre-existing secret %s is gone", existingSecret.Name) {
			continue
		}
		assert.Equalf(t, existingSecret.Data, found.Data, "data of pre-existing secret %s has changed", existingSecret.Name)
		delete(postDeletionSecretsByName, existingSecret.Name)
	}

	for _, expectedAfterDeletion := range c.ExpectedSecretsAfterDeletion {
		_, ok := postDeletionSecretsByName[expectedAfterDeletion]
		assert.Truef(t, ok, "secret %s missing, was expected to exist after deletion", expectedAfterDeletion)
		delete(postDeletionSecretsByName, expectedAfterDeletion)
	}

	assert.Empty(t, postDeletionSecretsByName, "newly created secrets remain after deletion")
}
