package argument

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Instance) GetInteger64(name string) int64 {
	v, e := i.flags.GetInt64(name)
	errors.PanicOnError(e)

	return v
}
