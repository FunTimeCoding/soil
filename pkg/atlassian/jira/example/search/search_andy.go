package search

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func searchAndy(
	j *jira.Client,
	p string,
	f *option.Format,
) {
	if true {
		console.Line("Search")
		issues := j.MustSearch(
			"project = %s AND status != %s",
			p,
			constant.JiraClosed,
		)
		console.Format("  Count: %d\n", len(issues))

		for _, i := range issues {
			console.Format("  Issue: %s\n", i.Format(f))
		}
	}

	if true {
		console.Line("SearchLimit")
		issues := j.MustSearchLimit(
			5,
			"project = %s AND status != %s",
			p,
			constant.JiraClosed,
		)
		console.Format("  Count: %d\n", len(issues))

		for _, i := range issues {
			console.Format("  Issue: %s\n", i.Format(f))
		}
	}
}
