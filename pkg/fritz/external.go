package fritz

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
)

func (c *Client) External() (*ExternalResponse, error) {
	body, e := c.wide(constant.ExternalAction)

	if e != nil {
		return nil, e
	}

	var result ExternalResponse

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
