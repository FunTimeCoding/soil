package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goalertmanagerd",
	"Alertmanager MCP server with multi-instance support",
	"goalertmanagerd",
).WithInstructions(
	"Alertmanager alert and silence management. Multi-instance - call list_instances and use_instance before querying.",
)

const (
	ListInstances = "list_instances"
	UseInstance   = "use_instance"
	ListAlerts    = "list_alerts"
	ListSilences  = "list_silences"
	CreateSilence = "create_silence"
	DeleteSilence = "delete_silence"
	GetStatus     = "get_status"
)
