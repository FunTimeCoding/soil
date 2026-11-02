package face

import "github.com/funtimecoding/soil/pkg/provision/types/change"

type DownstreamFunction func(c *change.Change) error
