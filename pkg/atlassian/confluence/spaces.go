package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

// Reference: https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-space/#api-spaces-get
func (c *Client) Spaces() ([]*space.Space, error) {
	l := c.basic.Base().Copy().Path(constant.ConfluenceSpace).Set(
		constant.ConfluenceStatus,
		constant.ConfluenceCurrentStatus,
	).String()
	var result []*response.Space

	for {
		var s *response.Spaces

		if e := c.basic.Get(l, &s); e != nil {
			return nil, e
		}

		result = append(result, s.Results...)

		if s.Links.Next == "" {
			break
		}

		l = c.basic.Next(s.Links.Next)
	}

	return space.Sort(space.NewSlice(result)), nil
}
