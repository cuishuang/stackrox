// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"

	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	// CreateTableTestg2grandchild1Stmt holds the create statement for table `testg2grandchild1`.
	CreateTableTestg2grandchild1Stmt = &postgres.CreateStmts{
		Table: `
               create table if not exists testg2grandchild1 (
                   Id varchar,
                   ParentId varchar,
                   ChildId varchar,
                   Val varchar,
                   serialized bytea,
                   PRIMARY KEY(Id),
                   CONSTRAINT fk_parent_table_0 FOREIGN KEY (ParentId) REFERENCES testggrandchild1(Id) ON DELETE CASCADE
               )
               `,
		Indexes:  []string{},
		Children: []*postgres.CreateStmts{},
	}

	// Testg2grandchild1Schema is the go schema for table `testg2grandchild1`.
	Testg2grandchild1Schema = func() *walker.Schema {
		schema := globaldb.GetSchemaForTable("testg2grandchild1")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.TestG2GrandChild1)(nil)), "testg2grandchild1").
			WithReference(Testggrandchild1Schema).
			WithReference(Testg3granchild1Schema)
		globaldb.RegisterTable(schema)
		return schema
	}()
)
