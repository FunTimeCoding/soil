package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"slices"
)

func (c *Client) Labeled() ([]*page.Page, error) {
	labels, e := c.Labels()

	if e != nil {
		return nil, e
	}

	var result []*page.Page
	var identifiers []string

	for _, l := range labels {
		if slices.Contains(c.labels, l.Name) {
			identifiers = append(identifiers, l.Identifier)
		}
	}

	for _, identifier := range identifiers {
		pages, f := c.PagesByLabel(identifier)

		if f != nil {
			return nil, f
		}

		result = append(result, pages...)
	}

	return result, nil
}
