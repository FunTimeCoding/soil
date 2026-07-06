package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
)

func scanClient(
	v *virtual_file_system.System,
	root string,
	path string,
	repo string,
) *Client {
	relative := strings.TrimPrefix(path, fmt.Sprintf("%s/", root))

	return &Client{
		Path:     fmt.Sprintf("pkg/%s", relative),
		Repo:     repo,
		Must:     hasMustFiles(v, path),
		Basic:    v.DirectoryExists(filepath.Join(path, "basic")),
		Constant: v.DirectoryExists(filepath.Join(path, "constant")),
		Example:  v.DirectoryExists(filepath.Join(path, "example")),
		Entity:   hasEntityPackages(v, path),
	}
}
