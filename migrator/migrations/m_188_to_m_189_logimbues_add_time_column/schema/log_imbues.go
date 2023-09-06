// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"
	"time"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
)

var (
	// CreateTableLogImbuesStmt holds the create statement for table `log_imbues`.
	CreateTableLogImbuesStmt = &postgres.CreateStmts{
		GormModel: (*LogImbues)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// LogImbuesSchema is the go schema for table `log_imbues`.
	LogImbuesSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.LogImbue)(nil)), "log_imbues")
		schema.ScopingResource = resources.Administration
		return schema
	}()
)

const (
	// LogImbuesTableName specifies the name of the table in postgres.
	LogImbuesTableName = "log_imbues"
)

// LogImbues holds the Gorm model for Postgres table `log_imbues`.
type LogImbues struct {
	ID         string     `gorm:"column:id;type:varchar;primaryKey"`
	Timestamp  *time.Time `gorm:"column:timestamp;type:timestamp"`
	Serialized []byte     `gorm:"column:serialized;type:bytea"`
}
