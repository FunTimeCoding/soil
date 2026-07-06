package raid_parser

import (
	"github.com/funtimecoding/soil/pkg/raid_parser/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Exists(constant.InsecureEnvironment),
	)
}
