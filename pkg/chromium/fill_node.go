package chromium

import "context"

func (c *Client) FillNode(
	x context.Context,
	backendNodeIdentifier int64,
	value string,
	direct bool,
) error {
	if direct {
		return c.fillNodeDirect(x, backendNodeIdentifier, value)
	}

	return c.fillNodeInsertText(x, backendNodeIdentifier, value)
}
