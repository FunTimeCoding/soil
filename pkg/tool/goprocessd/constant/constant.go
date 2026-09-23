package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goprocessd",
	"Process manager with environment reload",
	"goprocessd [flags]",
)

const HistoryCapacity = 200
const (
	HostEnvironment     = "GOPROCESS_HOST"
	PortEnvironment     = "GOPROCESS_PORT"
	InsecureEnvironment = "GOPROCESS_INSECURE"
	TokenEnvironment    = "GOPROCESS_TOKEN"
)
const (
	ProcessStatus     = "process_status"
	ProcessLog        = "process_log"
	ProcessRestart    = "process_restart"
	ProcessRestartAll = "process_restart_all"
	ProcessReload     = "process_reload"
)

const (
	UnknownProcess     = "unknown process %s"
	WaveAlreadyRunning = "restart wave already running"
	ProcfileScope      = "procfile"
	EnvironmentScope   = "environment"
)

const (
	GreenColor   = 32
	CyanColor    = 36
	MagentaColor = 35
	YellowColor  = 33
	BlueColor    = 34
	RedColor     = 31
)

var Colors = []int{
	GreenColor,
	CyanColor,
	MagentaColor,
	YellowColor,
	BlueColor,
	RedColor,
}
