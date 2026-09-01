package pull_request

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/check/pull_request/option"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
)

func Check(o *option.Request) {
	c := github.NewEnvironment()
	elements := collect(c, o)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	owner := c.MustUser().Name

	if o.Verbose {
		console.Format("Owner: %s\n", owner)
	}

	f := constant.Format.Copy().Tag(consoleConstant.TagTimestamp)

	for _, e := range elements {
		console.Line(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(monitorConstant.GoGitHubPullRequest.Plural)
	}
}
