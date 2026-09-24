package gohook

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"os"
	"path/filepath"
)

func Install(c *configuration.Configuration) {
	directory := c.HooksDirectory()

	if p := git.HooksPath(c.Root); p != "" &&
		system.AbsolutePath(filepath.Join(c.Root, p)) != directory {
		panic(fmt.Sprintf(constant.HooksPathRedirected, p, directory))
	}

	system.MakeDirectory(directory)

	for _, name := range c.HookNames() {
		path := filepath.Join(directory, name)
		errors.PanicOnError(
			os.WriteFile(
				path,
				[]byte(fmt.Sprintf(constant.Stub, name)),
				constant.StubMode,
			),
		)
		console.Format(constant.Installed, path)
	}
}
