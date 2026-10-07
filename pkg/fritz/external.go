package fritz

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
	"github.com/funtimecoding/soil/pkg/fritz/response"
)

func (c *Client) External() (*response.External, error) {
	body, e := c.wide(constant.ExternalAction)

	if e != nil {
		return nil, e
	}

	var result response.External

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
