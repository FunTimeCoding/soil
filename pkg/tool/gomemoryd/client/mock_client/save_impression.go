package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/impression_call"

func (c *Client) SaveImpression(
	content string,
	source string,
) {
	c.Impressions = append(
		c.Impressions,
		impression_call.Call{Content: content, Source: source},
	)
}
