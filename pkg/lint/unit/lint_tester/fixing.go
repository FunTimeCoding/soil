package lint_tester

import "github.com/funtimecoding/soil/pkg/lint/option"

func Fixing() *option.Lint {
	o := option.New("", false)
	o.Fix = true

	return o
}
