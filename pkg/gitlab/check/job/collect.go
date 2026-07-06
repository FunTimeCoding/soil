package job

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/monitor"
)

func collect(o *option.Job) []*job.Job {
	return monitor.OnlyConcerns(
		gitlab.NewEnvironment(gitlab.WithVerbose(o.Verbose)).Jobs(),
		o.All,
	)
}
