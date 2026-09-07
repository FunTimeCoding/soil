package process

import (
	"github.com/funtimecoding/soil/pkg/system/run/process"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/log"
	"sync"
	"time"
)

type Process struct {
	Name                string
	Command             string
	ColorIndex          int
	handle              *process.Process
	logger              *log.Logger
	stoppedBySupervisor bool
	waitError           error
	started             time.Time
	mutex               sync.Mutex
	condition           *sync.Cond
}
