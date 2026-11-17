package model_context_client

func (c *Client) ToolNames() []string {
	c.t.Helper()
	var result []string

	for _, tool := range c.ListTools() {
		result = append(result, tool.Name)
	}

	return result
}
