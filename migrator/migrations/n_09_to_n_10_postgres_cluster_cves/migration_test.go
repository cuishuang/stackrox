// Code originally generated by pg-bindings generator.

//go:build sql_integration
// +build sql_integration

package n9ton10

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	legacy "github.com/stackrox/rox/migrator/migrations/n_09_to_n_10_postgres_cluster_cves/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_09_to_n_10_postgres_cluster_cves/postgres"
	pghelper "github.com/stackrox/rox/migrator/migrations/postgreshelper"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/dackbox"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/rocksdb"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stackrox/rox/pkg/testutils/rocksdbtest"
	"github.com/stretchr/testify/suite"
)

func TestMigration(t *testing.T) {
	suite.Run(t, new(postgresMigrationSuite))
}

type postgresMigrationSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	ctx         context.Context

	legacyDB   *rocksdb.RocksDB
	postgresDB *pghelper.TestPostgres
}

var _ suite.TearDownTestSuite = (*postgresMigrationSuite)(nil)

func (s *postgresMigrationSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")
	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	var err error
	s.legacyDB, err = rocksdb.NewTemp(s.T().Name())
	s.NoError(err)

	s.Require().NoError(err)

	s.ctx = sac.WithAllAccess(context.Background())
	s.postgresDB = pghelper.ForT(s.T(), true)
}

func (s *postgresMigrationSuite) TearDownTest() {
	rocksdbtest.TearDownRocksDB(s.legacyDB)
	s.postgresDB.Teardown(s.T())
}

func (s *postgresMigrationSuite) TestClusterCVEMigration() {
	newStore := pgStore.New(s.postgresDB.Pool)
	dacky, err := dackbox.NewRocksDBDackBox(s.legacyDB, nil, []byte("graph"), []byte("dirty"), []byte("valid"))
	s.NoError(err)
	legacyStore := legacy.New(dacky, concurrency.NewKeyFence())

	// Prepare data and write to legacy DB
	var clusterCVEs []*storage.ClusterCVE
	for i := 0; i < 1; i++ {
		cve := &storage.CVE{}
		s.NoError(testutils.FullInit(cve, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		s.NoError(legacyStore.Upsert(s.ctx, cve))
		clusterCVEs = append(clusterCVEs, convertCVEToClusterCVEs(cve)...)
	}
	cve := &storage.CVE{
		Id:    "CVE-2019-02-14",
		Cvss:  1.3,
		Type:  storage.CVE_ISTIO_CVE,
		Types: []storage.CVE_CVEType{storage.CVE_K8S_CVE, storage.CVE_ISTIO_CVE},
	}
	s.NoError(legacyStore.Upsert(s.ctx, cve))
	clusterCVEs = append(clusterCVEs, convertCVEToClusterCVEs(cve)...)

	// Move
	s.NoError(move(s.postgresDB.GetGormDB(), s.postgresDB.Pool, legacyStore))

	// Verify
	count, err := newStore.Count(s.ctx)
	s.NoError(err)
	s.Equal(len(clusterCVEs), count)
	for _, clusterCVE := range clusterCVEs {
		fetched, exists, err := newStore.Get(s.ctx, clusterCVE.GetId())
		s.NoError(err)
		s.True(exists)
		s.Equal(clusterCVE, fetched)
	}
}
