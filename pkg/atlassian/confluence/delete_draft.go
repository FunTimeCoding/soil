package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (c *Client) DeleteDraft(pageIdentifier string) error {
	return c.basic.Delete(
		c.basic.Base().Copy().Path(
			"%s/%s",
			constant.ConfluencePage,
			pageIdentifier,
		).Set(constant.ConfluenceDraftParameter, "true").String(),
	)
}
