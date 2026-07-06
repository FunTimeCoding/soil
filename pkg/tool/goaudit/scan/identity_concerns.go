package scan

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func identityConcerns(
	v *virtual_file_system.System,
	path string,
	name string,
) []*concern.Concern {
	file := filepath.Join(path, "constant", "constant.go")

	if !v.Has(file) {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityMissingFileKey,
				IdentityMissingFileText,
				path,
			),
		}
	}

	return checkIdentitySource(v.ReadString(file), file, path, name)
}
