package mock_client

import "encoding/json"

func (c *Client) Evaluate(
	context string,
	expression string,
) (json.RawMessage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, e := c.find(context); e != nil {
		return nil, e
	}

	if result, okay := c.results[expression]; okay {
		return result, nil
	}

	return json.RawMessage(`{"type":"undefined"}`), nil
}
