package command_run

import (
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
)

type Run struct {
	Record *run.Run
	Hosts  []*host_result.Result
}
