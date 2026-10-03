package job

import (
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/check/job/option"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/monitor"
)

func collect(
	c *github.Client,
	o *option.Job,
) []*run.Run {
	return monitor.OnlyConcerns(
		c.MustRuns(true, o.Verbose && !o.Notation),
		o.All,
	)
}
