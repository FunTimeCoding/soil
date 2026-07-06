package argument

import "github.com/funtimecoding/soil/pkg/system"

func (i *Instance) RequiredInteger(name string) int {
	v := i.GetInteger(name)

	if v == 0 {
		system.Exitf(1, "flag empty: %s\n", name)
	}

	return v
}
