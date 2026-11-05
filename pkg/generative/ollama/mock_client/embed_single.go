package mock_client

import "context"

func (c *Client) EmbedSingle(
	_ context.Context,
	_ string,
) ([]float32, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Vector, nil
}
