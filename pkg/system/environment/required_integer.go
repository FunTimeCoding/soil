package environment

import "github.com/funtimecoding/soil/pkg/strings"

func RequiredInteger(name string) int {
	return strings.MustToInteger(Required(name))
}
