package mock_client

func (c *Client) Embedding(
	_ string,
	_ string,
) ([]float64, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.DoubleVector, nil
}
