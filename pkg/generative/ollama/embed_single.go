package ollama

import "context"

func (c *Client) EmbedSingle(
	x context.Context,
	v string,
) ([]float32, error) {
	result, e := c.Embed(x, []string{v})

	if e != nil {
		return nil, e
	}

	return result[0], nil
}
