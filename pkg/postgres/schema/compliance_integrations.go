// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableComplianceIntegrationsStmt holds the create statement for table `compliance_integrations`.
	CreateTableComplianceIntegrationsStmt = &postgres.CreateStmts{
		GormModel: (*ComplianceIntegrations)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ComplianceIntegrationsSchema is the go schema for table `compliance_integrations`.
	ComplianceIntegrationsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("compliance_integrations")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ComplianceIntegration)(nil)), "compliance_integrations")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Cluster":           ClustersSchema,
			"storage.NamespaceMetadata": NamespacesSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_COMPLIANCE_INTEGRATIONS, "complianceintegration", (*storage.ComplianceIntegration)(nil)))
		schema.SetSearchScope([]v1.SearchCategory{
			v1.SearchCategory_NAMESPACES,
			v1.SearchCategory_CLUSTERS,
		}...)
		schema.ScopingResource = resources.Compliance
		RegisterTable(schema, CreateTableComplianceIntegrationsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_COMPLIANCE_INTEGRATIONS, schema)
		return schema
	}()
)

const (
	// ComplianceIntegrationsTableName specifies the name of the table in postgres.
	ComplianceIntegrationsTableName = "compliance_integrations"
)

// ComplianceIntegrations holds the Gorm model for Postgres table `compliance_integrations`.
type ComplianceIntegrations struct {
	ID         string `gorm:"column:id;type:uuid;primaryKey"`
	Version    string `gorm:"column:version;type:varchar"`
	ClusterID  string `gorm:"column:clusterid;type:uuid;uniqueIndex:compliance_unique_indicator;index:complianceintegrations_sac_filter,type:hash"`
	Serialized []byte `gorm:"column:serialized;type:bytea"`
}
