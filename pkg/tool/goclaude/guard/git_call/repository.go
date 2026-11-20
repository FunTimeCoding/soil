package git_call

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"os"
	"path/filepath"
)

func (c *Call) Repository(directory string) string {
	if directory == "" || c.redirected {
		return ""
	}

	target := directory

	for _, d := range c.directories {
		if filepath.IsAbs(d) {
			target = d

			continue
		}

		target = filepath.Join(target, d)
	}

	resolved, e := filepath.EvalSymlinks(target)

	if e != nil {
		return ""
	}

	if i, f := os.Stat(resolved); f != nil || !i.IsDir() {
		return ""
	}

	return system.FindDirectoryUp(resolved, constant.GitMetadataDirectory)
}
