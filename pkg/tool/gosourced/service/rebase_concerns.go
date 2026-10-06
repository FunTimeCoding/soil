package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
)

func rebaseConcerns(
	root string,
	r *output.Results,
) []*concern.Concern {
	for _, c := range r.Entries {
		if c.Path != "" && !filepath.IsAbs(c.Path) {
			c.Path = filepath.Join(root, c.Path)
		}
	}

	return r.Entries
}
