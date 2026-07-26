package scan

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"path/filepath"
)

func IdentityConcerns(
	v *virtual_file_system.System,
	path string,
	name string,
) []*concern.Concern {
	file := filepath.Join(path, "constant", "constant.go")

	if !v.Has(file) {
		return []*concern.Concern{
			concern.NewPackage(
				constant.IdentityMissingFileKey,
				constant.IdentityMissingFileText,
				path,
			),
		}
	}

	return checkIdentitySource(v.ReadString(file), file, path, name)
}
