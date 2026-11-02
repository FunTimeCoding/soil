package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
)

type Client interface {
	Address() string
	Close()
	PostTrigger(c *change.Change) error
	Reachable(x context.Context) error
}
