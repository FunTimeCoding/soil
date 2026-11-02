package unit

import "github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"

func classify(s string) string {
	return lint_tester.Classify(
		s,
		[]string{".claude", ".claude-plugin", "doc", "pkg", "skills"},
	)
}
