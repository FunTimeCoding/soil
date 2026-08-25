package model_context_client

import "encoding/json"

func (c *Client) CallToolNotation(
	name string,
	arguments map[string]any,
) (map[string]any, error) {
	text, e := c.CallTool(name, arguments)

	if e != nil {
		return nil, e
	}

	var result map[string]any

	return result, json.Unmarshal([]byte(text), &result)
}
