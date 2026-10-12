package integration

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"path/filepath"
	"sync"
)

var binaryDirectory = sync.OnceValue(buildBinary)

func buildBinary() string {
	directory := system.TemporaryDirectory(constant.Identity.Name())
	r := run.New()
	r.Directory = git.FindDirectory()
	r.Start(
		soil.Go,
		soil.Build,
		soil.OutputArgument,
		filepath.Join(directory, constant.Identity.Name()),
		join.Slash(
			[]string{
				soil.CurrentDirectory,
				systemConstant.CommandPath,
				constant.Identity.Name(),
			},
		),
	)

	return directory
}
