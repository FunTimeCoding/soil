package argument

import "github.com/funtimecoding/soil/pkg/system"

func (i *Instance) RequiredPositional(
	number int,
	description string,
) string {
	if s := i.Argument(number); s != "" {
		return s
	}

	system.Exitf(1, "positional empty: %s\n", description)

	return ""
}
