package search

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
)

func searchOwn(
	j *jira.Client,
	p string,
) {
	if true {
		fmt.Println("SearchV3")
		issues := j.MustSearchV3(
			"project = %s AND status != %s",
			p,
			constant.Closed,
		)
		fmt.Printf("  Count: %d\n", len(issues))

		for _, i := range issues {
			fmt.Printf("  Issue: %s\n", i.Key)
		}
	}

	if true {
		fmt.Println("SearchLimitV3")
		issues := j.MustSearchLimitV3(
			5,
			"project = %s AND status != %s",
			p,
			constant.Closed,
		)
		fmt.Printf("  Count: %d\n", len(issues))

		for _, i := range issues {
			fmt.Printf("  Issue: %s\n", i.Key)
		}
	}
}
