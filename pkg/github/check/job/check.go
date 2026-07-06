package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/check/job/option"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/monitor"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
)

func Check(o *option.Job) {
	c := github.NewEnvironment()
	elements := collect(c, o)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	owner := c.MustUser().Name

	if o.Verbose {
		fmt.Printf("Owner: %s\n", owner)
	}

	f := constant.Format.Copy().Tag(tag.Timestamp)

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoGitHubJob.Plural)
	}
}
