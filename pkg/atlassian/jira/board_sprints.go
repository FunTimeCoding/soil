package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/integers"
)

func (c *Client) BoardSprints(identifier int) ([]*jira.Sprint, error) {
	v, _, e := c.client.Board.GetAllSprints(
		integers.ToString(identifier),
	)

	if e != nil {
		return nil, e
	}

	var result []*jira.Sprint

	for _, s := range v {
		result = append(result, &s)
	}

	return result, nil
}
