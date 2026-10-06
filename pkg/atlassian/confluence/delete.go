package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (c *Client) Delete(pageIdentifier string) error {
	return c.basic.Delete(
		c.basic.Base().Copy().Path(
			"%s/%s",
			constant.ConfluencePage,
			pageIdentifier,
		).String(),
	)
}
