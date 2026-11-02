package face

import "github.com/funtimecoding/soil/pkg/provision/types/change"

func (f DownstreamFunction) PostTrigger(c *change.Change) error {
	return f(c)
}
