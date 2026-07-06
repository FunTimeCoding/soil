package clean_job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func ProjectWay(
	g *gitlab.Client,
	p *project.Project,
	f *option.Format,
) {
	for _, j := range g.MustProjectJobs(p) {
		if j.Status != constant.Failed {
			continue
		}

		fmt.Printf("Job: %s\n", j.Format(f))

		if j.Trace != "" {
			fmt.Printf("  Trace: %s\n", j.Trace)
		}
	}
}
