package fritz

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
)

func (c *Client) Device() (*DeviceResponse, error) {
	body, e := c.call(
		constant.DevicePath,
		constant.DeviceService,
		constant.DeviceAction,
	)

	if e != nil {
		return nil, e
	}

	var result DeviceResponse

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
