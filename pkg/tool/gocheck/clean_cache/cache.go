package clean_cache

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/layer"
)

type Cache struct {
	handle    *index.Workspace
	layer     *layer.Layer
	faces     string
	comment   bool
	ours      []string
	listed    map[string][]string
	constants map[string]string
}
