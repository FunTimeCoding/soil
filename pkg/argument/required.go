package argument

import "github.com/funtimecoding/soil/pkg/system"

func (i *Instance) Required(name string) string {
	v := i.GetString(name)

	if v == "" {
		system.Exitf(1, "flag empty: %s\n", name)
	}

	return v
}
