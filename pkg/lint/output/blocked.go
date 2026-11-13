package output

import "github.com/funtimecoding/soil/pkg/lint/concern"

func Blocked(entries []*concern.Concern) bool {
	for _, c := range entries {
		if !c.Fixed && !c.Planned {
			return true
		}
	}

	return false
}
