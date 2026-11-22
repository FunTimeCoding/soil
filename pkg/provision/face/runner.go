package face

import (
	"github.com/funtimecoding/soil/pkg/provision/types/command_run"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

type Runner interface {
	Trigger(request trigger.Request) error
	HostCommand(
		command string,
		host string,
		source string,
	) (*command_run.Run, error)
}
