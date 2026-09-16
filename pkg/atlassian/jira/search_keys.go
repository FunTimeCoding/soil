package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) SearchKeys(
	query string,
	a ...any,
) ([]string, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	issues, e := c.searchV3Pages(query, constant.JiraKeyField, "")

	if e != nil {
		return nil, e
	}

	var result []string

	for _, i := range issues {
		result = append(result, i.Key)
	}

	return result, nil
}
