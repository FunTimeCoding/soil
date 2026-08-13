package ollama

func (c *Client) EmbedSingle(v string) ([]float32, error) {
	result, e := c.Embed([]string{v})

	if e != nil {
		return nil, e
	}

	return result[0], nil
}
