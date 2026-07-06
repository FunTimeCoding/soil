package bubbletea

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	library "github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"runtime"
)

func Run(
	m tea.Model,
	log bool,
) {
	if log {
		var path string

		if library.RunningFromBinary() {
			switch runtime.GOOS {
			case systemConstant.Darwin:
				path = join.Absolute(
					system.Home(),
					systemConstant.MacOSLibrary,
					systemConstant.MacOsLogs,
					constant.LogFile,
				)
			case systemConstant.Linux:
				path = join.Absolute(
					systemConstant.Temporary,
					constant.LogFile,
				)
			default:
				path = constant.LogFile
			}
		} else {
			system.EnsurePathExists(systemConstant.Temporary)
			path = join.Relative(systemConstant.Temporary, constant.LogFile)
		}

		f := Log(path)
		defer errors.LogClose(f)
	}

	RunProgram(tea.NewProgram(m))
}
