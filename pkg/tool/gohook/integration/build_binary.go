package integration

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	tool "github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"sync"
)

var binaryDirectory = sync.OnceValue(buildBinary)

func buildBinary() string {
	directory := system.TemporaryDirectory(tool.Identity.Name())
	r := run.New()
	r.Directory = git.FindDirectory()
	r.Start(
		constant.Go,
		constant.Build,
		constant.OutputArgument,
		filepath.Join(directory, tool.Identity.Name()),
		join.Slash(
			[]string{
				constant.CurrentDirectory,
				systemConstant.CommandPath,
				tool.Identity.Name(),
			},
		),
	)

	return directory
}
