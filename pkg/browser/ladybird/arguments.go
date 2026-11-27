package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strconv"
)

func (c *Client) arguments(port int) []string {
	result := []string{
		c.option.DriverPath,
		constant.LadybirdListenFlag,
		webConstant.Loopback,
		constant.LadybirdPortFlag,
		strconv.Itoa(port),
	}

	if c.option.Profile != "" {
		result = append(result, constant.LadybirdProfileFlag, c.option.Profile)
	}

	if c.option.Headless {
		result = append(result, constant.LadybirdHeadlessFlag)
	}

	return result
}
