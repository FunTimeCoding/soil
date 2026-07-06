package clean_job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"os"
)

func Check() {
	a := argument.NewSimple("clean-job")
	a.String(argument.Namespace, "", "Namespace")
	a.String(argument.Project, "", "Project")
	a.String(argument.Match, "", "Description match")
	a.ParseSimple()
	g := gitlab.NewEnvironment()
	f := constant.Format.Copy().Tag(tag.Dense)
	m := a.GetString(argument.Match)

	if m == "" {
		fmt.Printf(
			"--%s must match a runner description\n",
			argument.Match,
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
			a.Required(argument.Namespace),
			a.Required(argument.Project),
		)

		if false {
			PipelineWay(g, p, f)
		}

		if false {
			ProjectWay(g, p, f)
		}
	}
}
