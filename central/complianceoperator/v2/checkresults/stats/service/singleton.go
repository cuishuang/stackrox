package service

import (
	resultsDS "github.com/stackrox/rox/central/complianceoperator/v2/checkresults/datastore"
	complianceIntegrationDS "github.com/stackrox/rox/central/complianceoperator/v2/integration/datastore"
	profileDatastore "github.com/stackrox/rox/central/complianceoperator/v2/profiles/datastore"
	configDS "github.com/stackrox/rox/central/complianceoperator/v2/scanconfigurations/datastore"
	scanDS "github.com/stackrox/rox/central/complianceoperator/v2/scans/datastore"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	serviceInstance     Service
	serviceInstanceInit sync.Once
)

// Singleton returns the singleton instance of the compliance service.
func Singleton() Service {
	if !features.ComplianceEnhancements.Enabled() {
		return nil
	}

	serviceInstanceInit.Do(func() {
		serviceInstance = New(resultsDS.Singleton(), configDS.Singleton(), complianceIntegrationDS.Singleton(), profileDatastore.Singleton(), scanDS.Singleton())
	})
	return serviceInstance
}
