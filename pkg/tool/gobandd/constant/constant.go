package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gobandd",
	"Out-of-band power control over Intel AMT, multi-instance",
	"gobandd",
).WithInstructions(
	"Out-of-band power control for machines carrying Intel AMT. Call list_instances to see the management engines. shutdown and reboot ask the operating system through LMS by default; force cuts power below it.",
)

const (
	ListInstances = "list_instances"
	UseInstance   = "use_instance"
	Status        = "status"
	Start         = "start"
	Shutdown      = "shutdown"
	Reboot        = "reboot"
)

const (
	OperationStart    = "start"
	OperationShutdown = "shutdown"
	OperationReboot   = "reboot"
)

var Operations = []string{OperationStart, OperationShutdown, OperationReboot}

const (
	InstanceParameter = "instance"
	ForceParameter    = "force"

	InstanceKind = "instance"
)

const GracefulRefusal = "LMS on the host registers it; force cuts power instead"
