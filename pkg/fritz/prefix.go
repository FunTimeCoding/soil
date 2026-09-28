package fritz

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
)

func (c *Client) Prefix() (*PrefixResponse, error) {
	body, e := c.wide(constant.PrefixAction)

	if e != nil {
		return nil, e
	}

	var result PrefixResponse

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
