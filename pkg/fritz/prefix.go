package fritz

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
	"github.com/funtimecoding/soil/pkg/fritz/response"
)

func (c *Client) Prefix() (*response.Prefix, error) {
	body, e := c.wide(constant.PrefixAction)

	if e != nil {
		return nil, e
	}

	var result response.Prefix

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
