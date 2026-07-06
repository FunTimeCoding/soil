package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.UserKeyEnvironment),
		environment.Required(constant.TeamKeyEnvironment),
		environment.Required(constant.WebHostEnvironment),
	)
}
