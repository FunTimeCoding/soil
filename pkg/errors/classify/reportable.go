package classify

import "github.com/funtimecoding/soil/pkg/errors/validation"

func Reportable(e error) bool {
	return !validation.Is(e) && !expected(e)
}
