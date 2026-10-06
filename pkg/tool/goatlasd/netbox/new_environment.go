package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment() *client.ClientWithResponses {
	result, e := client.NewClientWithResponses(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		).String(),
		client.WithHTTPClient(web.StallClient()),
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)

	return result
}
