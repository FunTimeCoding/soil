package job

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/monitor"
)

func collect(o *option.Job) []*job.Job {
	c := gitlab.NewEnvironment(gitlab.WithVerbose(o.Verbose))
	defer c.Close()

	return monitor.OnlyConcerns(c.MustJobs(context.Background()), o.All)
}
