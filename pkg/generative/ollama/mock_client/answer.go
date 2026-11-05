package mock_client

func (c *Client) answer(prompt string) string {
	c.Prompts = append(c.Prompts, prompt)

	if c.Respond != nil {
		return c.Respond(prompt)
	}

	return c.Text
}
