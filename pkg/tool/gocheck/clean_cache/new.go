package clean_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"slices"
)

func New(
	handle *index.Workspace,
	l *layer.Layer,
	comment bool,
	ours []string,
) *Cache {
	return &Cache{
		handle:  handle,
		layer:   l,
		faces:   face.FromWorkspace(handle, nil).Fingerprint(),
		comment:   comment,
		ours:      slices.Sorted(slices.Values(ours)),
		listed:    make(map[string][]string),
		constants: make(map[string]string),
	}
}
