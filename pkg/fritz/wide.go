package fritz

import "github.com/funtimecoding/soil/pkg/fritz/constant"

func (c *Client) wide(action string) ([]byte, error) {
	body, e := c.call(constant.PPPPath, constant.PPPService, action)

	if e == nil {
		return body, nil
	}

	return c.call(constant.InternetPath, constant.InternetService, action)
}
