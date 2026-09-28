package band

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/band/constant"
)

func get[T any](
	c *Client,
	resource string,
) (*T, error) {
	body, e := c.call(constant.GetAction, resource, "", "")

	if e != nil {
		return nil, e
	}

	var result T

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
