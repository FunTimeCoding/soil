package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strconv"
)

func (c *Client) arguments(port int) []string {
	result := []string{
		c.option.DriverPath,
		constant.ListenFlag,
		webConstant.Loopback,
		constant.PortFlag,
		strconv.Itoa(port),
	}

	if c.option.Profile != "" {
		result = append(result, constant.ProfileFlag, c.option.Profile)
	}

	if c.option.Headless {
		result = append(result, constant.HeadlessFlag)
	}

	return result
}
