package terminal

import "github.com/funtimecoding/soil/pkg/system/environment"

func (t *Terminal) Required(name string) string {
	result := environment.Optional(name)

	if result == "" {
		t.Exitf("%s not set\n", name)
	}

	return result
}
