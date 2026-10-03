package clean_job

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"os"
)

func Check() {
	a := argument.NewSimple("clean-job")
	a.String(argumentConstant.Match, "", "Description match")
	a.ParseSimple()
	g := gitlab.NewEnvironment()
	f := constant.Format.Copy().Tag(consoleConstant.TagDense)
	m := a.GetString(argumentConstant.Match)

	if m == "" {
		console.Format(
			"--%s must match a runner description\n",
			argumentConstant.Match,
		)

		for _, r := range g.MustRunners(true) {
			console.Line(r.Format(f))
		}

		os.Exit(1)
	}

	r, found, e := g.FindRunnerByDescription(m)
	errors.PanicOnError(e)

	if !found {
		console.Line("No runner match")
		os.Exit(1)
	}

	RunnerWay(g, r, f)
}
