package inventory

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func DefaultPath() string {
	return filepath.Join(
		system.StorageDirectory(constant.StorageName, false),
		"gosourced.yaml",
	)
}
