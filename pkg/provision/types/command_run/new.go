package command_run

import (
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
)

func New(
	record *run.Run,
	hosts []*host_result.Result,
) *Run {
	return &Run{Record: record, Hosts: hosts}
}
