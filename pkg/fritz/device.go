package fritz

import (
	"encoding/xml"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
	"github.com/funtimecoding/soil/pkg/fritz/response"
)

func (c *Client) Device() (*response.Device, error) {
	body, e := c.call(
		constant.DevicePath,
		constant.DeviceService,
		constant.DeviceAction,
	)

	if e != nil {
		return nil, e
	}

	var result response.Device

	if f := xml.Unmarshal(body, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
