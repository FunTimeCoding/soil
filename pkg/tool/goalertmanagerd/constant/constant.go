package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
)

var Identity = identity.New(
	"goalertmanagerd",
	"Alertmanager MCP server with multi-instance support",
	"goalertmanagerd",
).WithInstructions(
	"Alertmanager alert and silence management.",
	paragraph.New(
		inventoryConstant.MultiInstance,
		inventoryConstant.MultiInstanceInstruction,
	),
)

const (
	ListAlerts    = "list_alerts"
	ListSilences  = "list_silences"
	CreateSilence = "create_silence"
	DeleteSilence = "delete_silence"
	GetStatus     = "get_status"
)
