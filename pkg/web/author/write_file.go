package author

func (c *Client) WriteFile(
	path string,
	content []byte,
) error {
	return c.client.Write(path, content, 0644)
}
