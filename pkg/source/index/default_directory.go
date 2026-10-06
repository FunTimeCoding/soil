package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func DefaultDirectory() string {
	return filepath.Join(
		system.StorageDirectory(constant.StorageName, true),
		constant.IndexDirectory,
	)
}
