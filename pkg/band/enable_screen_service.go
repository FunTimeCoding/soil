package band

import (
	"encoding/xml"
	"fmt"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
)

func (c *Client) EnableScreenService() error {
	body, e := c.call(
		constant.ScreenServiceStateAction,
		constant.ScreenServiceResource,
		constant.ScreenServiceSelectors,
		fmt.Sprintf(
			`<p:RequestStateChange_INPUT xmlns:p="%s"><p:RequestedState>%d</p:RequestedState></p:RequestStateChange_INPUT>`,
			constant.ScreenServiceResource,
			constant.ScreenServiceEnabled,
		),
	)

	if e != nil {
		return e
	}

	var result response.RequestStateChange

	if f := xml.Unmarshal(body, &result); f != nil {
		return f
	}

	if result.ReturnValue != 0 {
		return unexpected.Format(
			"screen service change refused: return value %d",
			result.ReturnValue,
		)
	}

	return nil
}
