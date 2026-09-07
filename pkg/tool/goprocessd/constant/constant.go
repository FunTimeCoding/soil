package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goprocessd",
	"Process manager with environment reload",
	"goprocessd [flags]",
)

const HistoryCapacity = 200
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

var Colors = []int{
	32, // green
	36, // cyan
	35, // magenta
	33, // yellow
	34, // blue
	31, // red
}
