package author

func (c *Client) ReadFile(path string) ([]byte, error) {
	return c.client.Read(path)
}
