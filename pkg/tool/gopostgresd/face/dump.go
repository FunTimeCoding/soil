package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
)

type Dump interface {
	Dump(x context.Context, i *instance.Instance, path string) error
	Count(x context.Context, path string) (int, error)
	Restore(
		x context.Context,
		i *instance.Instance,
		database string,
		path string,
		clean bool,
	) error
}
