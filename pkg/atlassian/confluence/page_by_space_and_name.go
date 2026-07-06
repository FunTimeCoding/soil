package confluence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) PageBySpaceAndName(
	spaceName string,
	name string,
) (*page.Page, error) {
	s, e := c.SpaceByName(spaceName)

	if e != nil {
		return nil, e
	}

	body, f := c.basic.GetV2(
		c.basic.Base().Copy().Path(constant.Page).Set(
			constant.BodyFormat,
			constant.StorageFormat,
		).Set(
			constant.SpaceIdentifier,
			s.Identifier,
		).Set(constant.Title, name).String(),
	)

	if f != nil {
		return nil, f
	}

	var result *response.Pages
	notation.MustDecode(body, &result, false)

	if len(result.Results) != 1 {
		return nil, fmt.Errorf(
			"expected 1 page named %s in space %s, got %d: %w",
			name,
			spaceName,
			len(result.Results),
			constant.ErrorNotFound,
		)
	}

	return page.New(result.Results[0], c.host), nil
}
