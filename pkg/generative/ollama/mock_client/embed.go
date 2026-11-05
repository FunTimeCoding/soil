package mock_client

import "context"

func (c *Client) Embed(
	_ context.Context,
	v []string,
) ([][]float32, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := make([][]float32, len(v))

	for i := range v {
		result[i] = c.Vector
	}

	return result, nil
}
