package technitium

import "encoding/json"

func (c *Client) get(
	path string,
	result any,
) error {
	payload, e := c.basic.Get(path)

	if e != nil {
		return e
	}

	return json.Unmarshal(payload, result)
}
