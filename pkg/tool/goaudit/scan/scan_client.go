package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/matrix"
	"path/filepath"
	"strings"
)

func scanClient(
	v *virtual_file_system.System,
	root string,
	path string,
	repo string,
) *matrix.Client {
	relative := strings.TrimPrefix(path, fmt.Sprintf("%s/", root))
	result := matrix.NewClient()
	result.Path = fmt.Sprintf("pkg/%s", relative)
	result.Repo = repo
	result.Must = hasMustFiles(v, path)
	result.Basic = v.DirectoryExists(filepath.Join(path, "basic"))
	result.Constant = v.DirectoryExists(
		filepath.Join(path, constant.ConstantDirectory),
	)
	result.Example = v.DirectoryExists(filepath.Join(path, "example"))
	result.Entity = hasEntityPackages(v, path)

	return result
}
