package mock_client

func (c *Client) SaveImpression(
	content string,
	source string,
) {
	c.Impressions = append(
		c.Impressions,
		ImpressionCall{Content: content, Source: source},
	)
}
