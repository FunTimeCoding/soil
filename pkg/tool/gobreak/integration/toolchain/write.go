package toolchain

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
)

func write(
	root string,
	name string,
	content string,
) {
	path := filepath.Join(root, name)
	system.MakeDirectory(filepath.Dir(path))
	system.WriteFile(path, []byte(content), constant.FileMode)
}
