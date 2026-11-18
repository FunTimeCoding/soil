package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/worker"
)

type Daemon struct {
	targets []*target.Target
	clock   *worker.Worker
}
