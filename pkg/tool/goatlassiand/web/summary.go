package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func summary(issues []*issue.Issue) []string {
	blocked := 0

	for _, i := range issues {
		if i.Blocked() {
			blocked++
		}
	}

	result := []string{fmt.Sprintf("%d on the plate", len(issues))}

	if blocked > 0 {
		result = append(result, fmt.Sprintf("%d blocked", blocked))
	}

	return result
}
