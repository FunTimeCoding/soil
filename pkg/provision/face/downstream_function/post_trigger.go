package downstream_function

import "github.com/funtimecoding/soil/pkg/provision/types/change"

func (f DownstreamFunction) PostTrigger(c *change.Change) error {
	return f(c)
}
