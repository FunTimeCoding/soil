package workspace

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
)

func New(indexDirectory string) *Service {
	return &Service{
		workspaces: cache.New(
			indexDirectory,
			face.Kind(),
			unclosed_resource.Kind(),
		),
		indexDirectory: indexDirectory,
	}
}
