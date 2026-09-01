package search

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/console"
)

func searchOwn(
	j *jira.Client,
	p string,
) {
	if true {
		console.Line("SearchV3")
		issues := j.MustSearchV3(
			"project = %s AND status != %s",
			p,
			constant.JiraClosed,
		)
		console.Format("  Count: %d\n", len(issues))

		for _, i := range issues {
			console.Format("  Issue: %s\n", i.Key)
		}
	}

	if true {
		console.Line("SearchLimitV3")
		issues := j.MustSearchLimitV3(
			5,
			"project = %s AND status != %s",
			p,
			constant.JiraClosed,
		)
		console.Format("  Count: %d\n", len(issues))

		for _, i := range issues {
			console.Format("  Issue: %s\n", i.Key)
		}
	}
}
