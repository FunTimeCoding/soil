package argument

import "github.com/funtimecoding/soil/pkg/strings/split"

func (i *Instance) Slice(name string) []string {
	v := i.GetString(name)

	if v == "" {
		return []string{}
	}

	return split.Comma(v)
}
