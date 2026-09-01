package clean_job

import (
	"github.com/funtimecoding/soil/pkg/console"
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
		if j.Status != constant.JobFail {
			continue
		}

		console.Format("Job: %s\n", j.Format(f))

		if j.Trace != "" {
			console.Format("  Trace: %s\n", j.Trace)
		}
	}
}
