package lint_tester

import "github.com/funtimecoding/soil/pkg/lint/pointer"

func Classify(s string) string {
	return string(
		pointer.Classify(
			s,
			[]string{".claude", ".claude-plugin", "doc", "pkg", "skills"},
		),
	)
}
