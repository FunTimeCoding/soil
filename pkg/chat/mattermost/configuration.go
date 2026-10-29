package mattermost

import "github.com/funtimecoding/soil/pkg/chat/constant"

func (c *Client) Configuration() (map[string]string, error) {
	result, r, e := c.client.GetClientConfig(
		c.context,
		constant.MattermostEmptyEntityTag,
	)

	return result, wrapError(r, e)
}
