package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goband",
	"Out-of-band power control CLI for gobandd",
	"goband <command>",
)

const (
	HostEnvironment     = "GOBAND_HOST"
	PortEnvironment     = "GOBAND_PORT"
	InsecureEnvironment = "GOBAND_INSECURE"
	TokenEnvironment    = "GOBAND_TOKEN"

	InstanceArgument = "instance"
	InstanceUsage    = "Management engine name (optional when only one is configured)"
	YesArgument      = "yes"
	ForceArgument    = "force"
)

const (
	StartDescription         = "Would power the machine on - the firmware only offers this while it is off. Pass --yes to execute."
	ShutdownDescription      = "Would ask the operating system to shut down through LMS. Pass --yes to execute."
	ShutdownForceDescription = "Would cut power below the operating system - it gets no chance to shut down. Pass --yes to execute."
	RebootDescription        = "Would ask the operating system to restart through LMS. Pass --yes to execute."
	RebootForceDescription   = "Would power cycle below the operating system - it gets no chance to shut down. Pass --yes to execute."
	GracefulFollowUp         = "Poll status - if the machine never goes down, the operating system is not answering and --force is the remaining option."
)
