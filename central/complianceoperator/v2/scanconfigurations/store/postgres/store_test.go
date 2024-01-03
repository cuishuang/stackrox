// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type ComplianceOperatorScanConfigurationV2StoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceOperatorScanConfigurationV2Store(t *testing.T) {
	suite.Run(t, new(ComplianceOperatorScanConfigurationV2StoreSuite))
}

func (s *ComplianceOperatorScanConfigurationV2StoreSuite) SetupSuite() {

	s.T().Setenv(features.ComplianceEnhancements.EnvVar(), "true")
	if !features.ComplianceEnhancements.Enabled() {
		s.T().Skip("Skip postgres store tests because feature flag is off")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceOperatorScanConfigurationV2StoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_operator_scan_configuration_v2 CASCADE")
	s.T().Log("compliance_operator_scan_configuration_v2", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *ComplianceOperatorScanConfigurationV2StoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceOperatorScanConfigurationV2StoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceOperatorScanConfigurationV2 := &storage.ComplianceOperatorScanConfigurationV2{}
	s.NoError(testutils.FullInit(complianceOperatorScanConfigurationV2, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceOperatorScanConfigurationV2, exists, err := store.Get(ctx, complianceOperatorScanConfigurationV2.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorScanConfigurationV2)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceOperatorScanConfigurationV2))
	foundComplianceOperatorScanConfigurationV2, exists, err = store.Get(ctx, complianceOperatorScanConfigurationV2.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceOperatorScanConfigurationV2, foundComplianceOperatorScanConfigurationV2)

	complianceOperatorScanConfigurationV2Count, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, complianceOperatorScanConfigurationV2Count)
	complianceOperatorScanConfigurationV2Count, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(complianceOperatorScanConfigurationV2Count)

	complianceOperatorScanConfigurationV2Exists, err := store.Exists(ctx, complianceOperatorScanConfigurationV2.GetId())
	s.NoError(err)
	s.True(complianceOperatorScanConfigurationV2Exists)
	s.NoError(store.Upsert(ctx, complianceOperatorScanConfigurationV2))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceOperatorScanConfigurationV2), sac.ErrResourceAccessDenied)

	foundComplianceOperatorScanConfigurationV2, exists, err = store.Get(ctx, complianceOperatorScanConfigurationV2.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceOperatorScanConfigurationV2, foundComplianceOperatorScanConfigurationV2)

	s.NoError(store.Delete(ctx, complianceOperatorScanConfigurationV2.GetId()))
	foundComplianceOperatorScanConfigurationV2, exists, err = store.Get(ctx, complianceOperatorScanConfigurationV2.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorScanConfigurationV2)
	s.NoError(store.Delete(withNoAccessCtx, complianceOperatorScanConfigurationV2.GetId()))

	var complianceOperatorScanConfigurationV2s []*storage.ComplianceOperatorScanConfigurationV2
	var complianceOperatorScanConfigurationV2IDs []string
	for i := 0; i < 200; i++ {
		complianceOperatorScanConfigurationV2 := &storage.ComplianceOperatorScanConfigurationV2{}
		s.NoError(testutils.FullInit(complianceOperatorScanConfigurationV2, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceOperatorScanConfigurationV2s = append(complianceOperatorScanConfigurationV2s, complianceOperatorScanConfigurationV2)
		complianceOperatorScanConfigurationV2IDs = append(complianceOperatorScanConfigurationV2IDs, complianceOperatorScanConfigurationV2.GetId())
	}

	s.NoError(store.UpsertMany(ctx, complianceOperatorScanConfigurationV2s))

	complianceOperatorScanConfigurationV2Count, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, complianceOperatorScanConfigurationV2Count)

	s.NoError(store.DeleteMany(ctx, complianceOperatorScanConfigurationV2IDs))

	complianceOperatorScanConfigurationV2Count, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, complianceOperatorScanConfigurationV2Count)
}
