// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"fmt"
	"reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableProcessIndicatorsStmt holds the create statement for table `process_indicators`.
	CreateTableProcessIndicatorsStmt = &postgres.CreateStmts{
		GormModel: (*ProcessIndicators)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// ProcessIndicatorsSchema is the go schema for table `process_indicators`.
	ProcessIndicatorsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("process_indicators")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.ProcessIndicator)(nil)), "process_indicators")
		referencedSchemas := map[string]*walker.Schema{
			"storage.Deployment": DeploymentsSchema,
		}

		schema.ResolveReferences(func(messageTypeName string) *walker.Schema {
			return referencedSchemas[fmt.Sprintf("storage.%s", messageTypeName)]
		})
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_PROCESS_INDICATORS, "processindicator", (*storage.ProcessIndicator)(nil)))
		RegisterTable(schema, CreateTableProcessIndicatorsStmt)
		return schema
	}()
)

const (
	ProcessIndicatorsTableName = "process_indicators"
)

// ProcessIndicators holds the Gorm model for Postgres table `process_indicators`.
type ProcessIndicators struct {
	Id                 string `gorm:"column:id;type:varchar;primaryKey"`
	DeploymentId       string `gorm:"column:deploymentid;type:varchar;index:processindicators_deploymentid,type:hash"`
	ContainerName      string `gorm:"column:containername;type:varchar"`
	PodId              string `gorm:"column:podid;type:varchar"`
	PodUid             string `gorm:"column:poduid;type:varchar;index:processindicators_poduid,type:hash"`
	SignalContainerId  string `gorm:"column:signal_containerid;type:varchar"`
	SignalName         string `gorm:"column:signal_name;type:varchar"`
	SignalArgs         string `gorm:"column:signal_args;type:varchar"`
	SignalExecFilePath string `gorm:"column:signal_execfilepath;type:varchar"`
	SignalUid          uint32 `gorm:"column:signal_uid;type:integer"`
	ClusterId          string `gorm:"column:clusterid;type:varchar"`
	Namespace          string `gorm:"column:namespace;type:varchar"`
	Serialized         []byte `gorm:"column:serialized;type:bytea"`
}
