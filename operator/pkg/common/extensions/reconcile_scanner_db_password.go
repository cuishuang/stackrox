package extensions

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/operator/pkg/types"
	"github.com/stackrox/rox/pkg/renderer"
	ctrlClient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	scannerDBPasswordKey          = `password`
	scannerDBPasswordResourceName = "scanner-db-password"
)

// ScannerBearingCustomResource interface exposes details about the Scanner resource from the kubernetes object.
type ScannerBearingCustomResource interface {
	types.K8sObject
	IsScannerEnabled() bool
}

// reconcileScannerDBPasswordConfig represents the config for scanner db password reconciliation
type reconcileScannerDBPasswordConfig struct {
	PasswordResourceName string
}

// ReconcileScannerDBPassword reconciles a scanner db password
func ReconcileScannerDBPassword(ctx context.Context, obj ScannerBearingCustomResource, client ctrlClient.Client, direct ctrlClient.Reader) error {
	return reconcileScannerDBPassword(ctx, obj, client, direct, reconcileScannerDBPasswordConfig{
		PasswordResourceName: scannerDBPasswordResourceName,
	})
}

func reconcileScannerDBPassword(ctx context.Context, obj ScannerBearingCustomResource, client ctrlClient.Client, direct ctrlClient.Reader, config reconcileScannerDBPasswordConfig) error {
	run := &reconcileScannerDBPasswordExtensionRun{
		// This is using OwnershipStrategyOwnerReference, so the secret will be garbage-collected
		// when SecuredCluster CR is deleted. This is because ScannerV2 database uses an ephemeral
		// storage, so keeping the password secret around after the CR is deleted is not useful.
		SecretReconciliator:  NewSecretReconciliator(client, direct, obj, OwnershipStrategyOwnerReference),
		obj:                  obj,
		passwordResourceName: config.PasswordResourceName,
	}
	return run.Execute(ctx)
}

type reconcileScannerDBPasswordExtensionRun struct {
	*SecretReconciliator
	obj                  ScannerBearingCustomResource
	passwordResourceName string
}

func (r *reconcileScannerDBPasswordExtensionRun) Execute(ctx context.Context) error {
	// Delete any scanner-db password only if the CR is being deleted, or scanner is not enabled.
	shouldExist := r.obj.GetDeletionTimestamp() == nil && r.obj.IsScannerEnabled()

	if err := r.reconcilePasswordSecret(ctx, shouldExist); err != nil {
		return errors.Wrapf(err, "reconciling %q secret", r.passwordResourceName)
	}

	return nil
}

func (r *reconcileScannerDBPasswordExtensionRun) reconcilePasswordSecret(ctx context.Context, shouldExist bool) error {
	if shouldExist {
		return r.EnsureSecret(ctx, r.passwordResourceName, r.validateScannerDBPasswordData, r.generateScannerDBPasswordData)
	}
	return r.DeleteSecret(ctx, r.passwordResourceName)
}

func (r *reconcileScannerDBPasswordExtensionRun) validateScannerDBPasswordData(data types.SecretDataMap, _ bool) error {
	if len(data[scannerDBPasswordKey]) == 0 {
		return errors.Errorf("%s secret must contain a non-empty %q entry", r.passwordResourceName, scannerDBPasswordKey)
	}
	return nil
}

func (r *reconcileScannerDBPasswordExtensionRun) generateScannerDBPasswordData(_ types.SecretDataMap) (types.SecretDataMap, error) {
	data := types.SecretDataMap{
		scannerDBPasswordKey: []byte(renderer.CreatePassword()),
	}
	return data, nil
}
