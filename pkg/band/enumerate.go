package band

import (
	"encoding/xml"
	"fmt"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
)

func enumerate[T any](
	c *Client,
	resource string,
) (*T, error) {
	opened, e := c.call(
		constant.EnumerateAction,
		resource,
		"",
		fmt.Sprintf(`<Enumerate xmlns="%s"/>`, constant.EnumerationNamespace),
	)

	if e != nil {
		return nil, e
	}

	var context response.Enumerate

	if f := xml.Unmarshal(opened, &context); f != nil {
		return nil, f
	}

	pulled, g := c.call(
		constant.PullAction,
		resource,
		"",
		fmt.Sprintf(
			`<Pull xmlns="%s"><EnumerationContext>%s</EnumerationContext></Pull>`,
			constant.EnumerationNamespace,
			context.Context,
		),
	)

	if g != nil {
		return nil, g
	}

	var result T

	if h := xml.Unmarshal(pulled, &result); h != nil {
		return nil, h
	}

	return &result, nil
}
