package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
)

var Identity = identity.New(
	"gopostgresd",
	"Multi-instance PostgreSQL access bridge",
	"gopostgresd",
).WithInstructions(
	"PostgreSQL access.",
	paragraph.New(
		inventoryConstant.MultiInstance,
		inventoryConstant.MultiInstanceInstruction,
	),
)

const (
	Query         = "query"
	Explain       = "explain"
	ListSchemas   = "list_schemas"
	ListTables    = "list_tables"
	DescribeTable = "describe_table"
	ListIndexes   = "list_indexes"
	TableSizes    = "table_sizes"
)

const (
	DefaultSchema = "public"
	InstancePeer  = "postgres"
)
