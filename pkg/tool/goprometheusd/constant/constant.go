package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goprometheusd",
	"Prometheus MCP server with multi-instance support",
	"goprometheusd",
).WithInstructions(
	"Prometheus query interface. Multi-instance - call list_instances and use_instance before querying.",
)

const (
	ListInstances = "list_instances"
	UseInstance   = "use_instance"
	Query         = "query"
	QueryRange    = "query_range"
	Series        = "series"
	LabelValues   = "label_values"
	LabelNames    = "label_names"
	QueryRules    = "query_rules"
)
