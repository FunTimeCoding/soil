package lint_tester

import "github.com/funtimecoding/soil/pkg/lint/pointer"

func Classify(
	s string,
	roots []string,
) string {
	return string(pointer.Classify(s, roots))
}
