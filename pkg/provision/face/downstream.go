package face

import "github.com/funtimecoding/soil/pkg/provision/types/change"

type Downstream interface {
	Trigger(c *change.Change) error
}
