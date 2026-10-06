package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) PageBySpaceAndName(
	spaceName string,
	name string,
) (*page.Page, error) {
	s, e := c.SpaceByName(spaceName)

	if e != nil {
		return nil, e
	}

	var result *response.Pages

	if f := c.basic.Get(
		c.basic.Base().Copy().Path(constant.ConfluencePage).Set(
			constant.ConfluenceBodyFormat,
			constant.ConfluenceStorageFormat,
		).Set(
			constant.ConfluenceSpaceIdentifier,
			s.Identifier,
		).Set(constant.ConfluenceTitle, name).String(),
		&result,
	); f != nil {
		return nil, f
	}

	if len(result.Results) == 0 {
		return nil, not_found.Format(
			"page not found: %s (space %s)",
			name,
			spaceName,
		)
	}

	if len(result.Results) > 1 {
		return nil, ambiguous.Format(
			"expected 1 page named %s in space %s, got %d",
			name,
			spaceName,
			len(result.Results),
		)
	}

	return page.New(result.Results[0], c.host), nil
}
