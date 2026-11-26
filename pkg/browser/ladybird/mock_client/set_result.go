package mock_client

import "encoding/json"

func (c *Client) SetResult(
	expression string,
	result string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.results[expression] = json.RawMessage(result)
}
