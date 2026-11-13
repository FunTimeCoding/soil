package gobreak

import (
	"fmt"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_result"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
)

func Results(path string) ([]*break_result.Result, int, error) {
	file := system.AbsolutePath(path)
	root := system.FindDirectoryUp(filepath.Dir(file), gitConstant.Directory)

	if root == "" {
		return nil, 0, fmt.Errorf(constant.NoRepository, file)
	}

	directory := system.TemporaryDirectory(constant.TemporaryPrefix)
	defer system.Remove(directory)

	return breaks(root, break_file.Load(file), directory)
}
