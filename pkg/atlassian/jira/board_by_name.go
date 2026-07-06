package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
)

func (c *Client) BoardByName(name string) (*jira.Board, error) {
	boards, e := c.Boards()

	if e != nil {
		return nil, e
	}

	for _, b := range boards {
		if b.Name == name {
			return b, nil
		}
	}

	return nil, fmt.Errorf(
		"board not found: %s: %w",
		name,
		constant.ErrorNotFound,
	)
}
