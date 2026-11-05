package github

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/github/refusal"
)

func (c *Client) SearchIssue(
	query string,
	a ...any,
) ([]*issue.Issue, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, _, e := c.client.Search.Issues(c.context, query, nil)

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return issue.NewSlice(result.Issues), nil
}
