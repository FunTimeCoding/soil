package gmail

import (
	"github.com/funtimecoding/soil/pkg/gmail/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func NewEnvironment() *Client {
	return New(
		environment.Fallback(
			constant.DirectoryEnvironment,
			join.Absolute(system.Home(), constant.DefaultDirectory),
		),
	)
}
