package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"

func (c *Client) DeleteDraft(pageIdentifier string) error {
	_, e := c.basic.DeleteV2(
		c.basic.Base().Copy().Path(
			"%s/%s",
			constant.Page,
			pageIdentifier,
		).Set(constant.Draft, "true").String(),
	)

	return e
}
