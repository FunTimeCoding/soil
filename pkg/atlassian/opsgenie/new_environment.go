package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.OpsgenieUserKeyEnvironment),
		environment.Required(constant.OpsgenieTeamKeyEnvironment),
		environment.Required(constant.OpsgenieWebHostEnvironment),
	)
}
