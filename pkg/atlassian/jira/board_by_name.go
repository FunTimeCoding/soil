package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
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

	return nil, not_found.New("board", name)
}
