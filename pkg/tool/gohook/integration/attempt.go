package integration

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	system "github.com/funtimecoding/soil/pkg/system/constant"
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
		system.PathEnvironment,
		join.Empty(
			binaryDirectory(),
			string(os.PathListSeparator),
			os.Getenv(system.PathEnvironment),
		),
	)

	for k, v := range environment {
		r.Environment(k, v)
	}

	r.Start(append([]string{constant.Command}, arguments...)...)

	return r
}
