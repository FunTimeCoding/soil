package lint

import "github.com/funtimecoding/soil/pkg/strings/join"

func PatternDetail(patterns []string) string {
	if len(patterns) == 0 {
		return ""
	}

	return join.Space("patterns", join.Space(patterns...))
}
