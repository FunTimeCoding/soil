package loki

func (c *Client) Series(series string) (string, error) {
	return c.basic.Series(series)
}
