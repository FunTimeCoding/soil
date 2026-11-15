package fix_tester

import "github.com/funtimecoding/soil/pkg/lint/concern"

func Applied(entries []*concern.Concern) []*concern.Concern {
	var r []*concern.Concern

	for _, c := range entries {
		if c.Fixed {
			r = append(r, c)
		}
	}

	return r
}
