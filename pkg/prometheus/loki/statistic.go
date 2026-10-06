package loki

func (c *Client) Statistic(query string) (string, error) {
	return c.basic.Statistic(query)
}
