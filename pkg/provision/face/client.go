package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
)

type Client interface {
	Address() string
	Close()
	PostTrigger(c *change.Change) error
	PostHostCommand(
		x context.Context,
		command string,
		host string,
		source string,
	) (client.HostCommandRun, error)
	Reachable(x context.Context) error
}
