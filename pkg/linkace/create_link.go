package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) CreateLink(
	l string,
	title string,
	listIdentifier int,
	tags []string,
) (*link.Link, error) {
	body := map[string]any{
		"url":         l,
		"title":       title,
		"description": "",
		"lists":       []int{listIdentifier},
		"tags":        tags,
		"visibility":  3,
	}
	var r response.Link

	if e := c.basic.Post("links", body, &r); e != nil {
		return nil, e
	}

	return c.UpdateLink(
		r.Identifier,
		map[string]any{
			"url":   l,
			"lists": []int{listIdentifier},
			"tags":  tags,
		},
	)
}
