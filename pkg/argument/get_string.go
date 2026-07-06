package argument

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Instance) GetString(name string) string {
	v, e := i.flags.GetString(name)
	errors.PanicOnError(e)

	return v
}
