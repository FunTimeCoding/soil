package system

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
)

func StorageDirectory(
	name string,
	create bool,
) string {
	result := filepath.Join(Home(), constant.StoragePath, name)

	if create {
		MakeDirectory(result)
	}

	return result
}
