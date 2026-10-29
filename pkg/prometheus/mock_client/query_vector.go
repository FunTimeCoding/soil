package mock_client

func (c *Client) QueryVector(_ string) (float64, error) {
	return 0, c.Failure
}
