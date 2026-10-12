package integration

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"os"
)

func attempt(
	directory string,
	environment map[string]string,
	arguments ...string,
) *run.Run {
	r := run.New()
	r.Panic = false
	r.Directory = directory
	r.Environment(
		systemConstant.PathEnvironment,
		join.Empty(
			binaryDirectory(),
			string(os.PathListSeparator),
			os.Getenv(systemConstant.PathEnvironment),
		),
	)

	for k, v := range environment {
		r.Environment(k, v)
	}

	r.Start(append([]string{gitConstant.Command}, arguments...)...)

	return r
}
