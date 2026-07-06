package argument

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Instance) GetInteger(name string) int {
	v, e := i.flags.GetInt(name)
	errors.PanicOnError(e)

	return v
}
