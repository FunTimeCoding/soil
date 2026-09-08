package utilization

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func fileRaw() string {
	path := filepath.Join(
		system.Home(),
		constant.ClaudeDirectory,
		constant.ClaudeCredentialFile,
	)

	if !system.FileExists(path) {
		return ""
	}

	return system.ReadFileUnsafe(path)
}
