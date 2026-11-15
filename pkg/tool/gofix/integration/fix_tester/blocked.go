package fix_tester

import "github.com/funtimecoding/soil/pkg/lint/concern"

func Blocked(entries []*concern.Concern) []*concern.Concern {
	var r []*concern.Concern

	for _, c := range entries {
		if !c.Fixed && !c.Planned {
			r = append(r, c)
		}
	}

	return r
}
