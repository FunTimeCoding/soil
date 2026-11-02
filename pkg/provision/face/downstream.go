package face

import "github.com/funtimecoding/soil/pkg/provision/types/change"

type Downstream interface {
	PostTrigger(c *change.Change) error
}
