package argument

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (i *Instance) GetDuration(name string) time.Duration {
	v, e := i.flags.GetDuration(name)
	errors.PanicOnError(e)

	return v
}
