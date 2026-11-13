package repository

import (
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func New(
	root string,
	files *virtual_file_system.System,
) *Repository {
	return NewWithForest(root, files, warm_forest.New())
}
