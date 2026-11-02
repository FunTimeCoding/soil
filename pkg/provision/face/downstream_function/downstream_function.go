package downstream_function

import "github.com/funtimecoding/soil/pkg/provision/types/change"

type DownstreamFunction func(c *change.Change) error
