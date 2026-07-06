package clean_job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func PipelineWay(
	g *gitlab.Client,
	p *project.Project,
	f *option.Format,
) {
	for _, i := range g.MustPipelines(p.Identifier) {
		fmt.Printf("Pipeline: %+v\n", i.ID)

		for _, j := range g.MustPipelineJobs(p.Identifier, i.ID) {
			fmt.Printf("  Job: %s\n", j.Format(f))
		}
	}
}
