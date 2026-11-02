package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
)

var Identity = identity.New(
	"goprometheusd",
	"Prometheus MCP server with multi-instance support",
	"goprometheusd",
).WithInstructions(
	"Prometheus query interface.",
	paragraph.New(
		inventoryConstant.MultiInstance,
		inventoryConstant.MultiInstanceInstruction,
	),
)

const (
	Query       = "query"
	QueryRange  = "query_range"
	Series      = "series"
	LabelValues = "label_values"
	LabelNames  = "label_names"
	QueryRules  = "query_rules"
)

const InstancePeer = "prometheus"
