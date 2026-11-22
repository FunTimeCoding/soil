package face

import (
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

type Runner interface {
	Trigger(request trigger.Request) error
	HostCommand(
		command string,
		host string,
		source string,
	) (*run.Run, error)
}
