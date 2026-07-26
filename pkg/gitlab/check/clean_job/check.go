package clean_job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"os"
)

func Check() {
	a := argument.NewSimple("clean-job")
	a.String(argumentConstant.Namespace, "", "Namespace")
	a.String(argumentConstant.Project, "", "Project")
	a.String(argumentConstant.Match, "", "Description match")
	a.ParseSimple()
	g := gitlab.NewEnvironment()
	f := constant.Format.Copy().Tag(tag.Dense)
	m := a.GetString(argumentConstant.Match)

	if m == "" {
		fmt.Printf(
			"--%s must match a runner description\n",
			argumentConstant.Match,
		)

		for _, r := range g.MustRunners(true) {
			fmt.Println(r.Format(f))
		}

		os.Exit(1)
	}

	r := g.RunnerByDescriptionMatch(m)

	if r == nil {
		fmt.Println("No runner match")
		os.Exit(1)
	}

	RunnerWay(g, r, f)

	if false {
		p := g.MustProjectByName(
			a.Required(argumentConstant.Namespace),
			a.Required(argumentConstant.Project),
		)

		if false {
			PipelineWay(g, p, f)
		}

		if false {
			ProjectWay(g, p, f)
		}
	}
}
