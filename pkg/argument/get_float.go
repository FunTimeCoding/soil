package argument

import "github.com/funtimecoding/soil/pkg/errors"

func (i *Instance) GetFloat(name string) float64 {
	v, e := i.flags.GetFloat64(name)
	errors.PanicOnError(e)

	return v
}
