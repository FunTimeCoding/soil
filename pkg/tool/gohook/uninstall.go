package gohook

import (
	"github.com/funtimecoding/soil/pkg/console"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"strings"
)

func Uninstall(c *configuration.Configuration) {
	directory := c.HooksDirectory()

	for _, name := range git.Hooks {
		path := filepath.Join(directory, name)

		if !system.FileExists(path) {
			continue
		}

		if !strings.Contains(system.ReadFileUnsafe(path), constant.StubMarker) {
			console.Format(constant.Kept, path)

			continue
		}

		system.RemoveFile(path)
		console.Format(constant.Removed, path)
	}
}
