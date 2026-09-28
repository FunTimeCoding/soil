package fritz

import "github.com/funtimecoding/soil/pkg/fritz/constant"

// The connection service depends on the uplink type: PPP for
// DSL, IP for cable and fiber. Try PPP first on this DSL box,
// fall back to IP.
func (c *Client) wide(action string) ([]byte, error) {
	body, e := c.call(constant.PPPPath, constant.PPPService, action)

	if e == nil {
		return body, nil
	}

	return c.call(constant.InternetPath, constant.InternetService, action)
}
