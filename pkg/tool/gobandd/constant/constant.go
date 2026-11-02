package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

var Identity = identity.New(
	"gobandd",
	"Out-of-band power control over Intel AMT, multi-instance",
	"gobandd",
).WithInstructions(
	"Out-of-band power control for machines carrying Intel AMT. shutdown and reboot ask the operating system through LMS by default; force cuts power below it.",
	paragraph.New(
		inventoryConstant.MultiInstance,
		inventoryConstant.MultiInstanceInstruction,
	),
)

const (
	Status   = "status"
	Start    = "start"
	Shutdown = "shutdown"
	Reboot   = "reboot"

	EnableRedirection = "enable_redirection"
)

const (
	ListenerWrite      = "listener"
	ScreenServiceWrite = "screen service"
)

const (
	HostEnvironment     = "GOBAND_HOST"
	PortEnvironment     = "GOBAND_PORT"
	InsecureEnvironment = "GOBAND_INSECURE"
	TokenEnvironment    = "GOBAND_TOKEN"
)

const (
	OperationStart    = "start"
	OperationShutdown = "shutdown"
	OperationReboot   = "reboot"
)

var Operations = []string{OperationStart, OperationShutdown, OperationReboot}
var IntentStatus = map[string]string{
	OperationStart:    netbox.StatusActive,
	OperationShutdown: netbox.StatusOffline,
}

const ForceParameter = "force"
const GracefulRefusal = "LMS on the host registers it; force cuts power instead"
