package process

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	processConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewOptional() *Collector {
	place := environment.Optional(constant.PlaceEnvironment)
	token := environment.Optional(processConstant.TokenEnvironment)

	if place == "" || token == "" {
		return nil
	}

	c, e := client.NewClientWithResponses(
		locator.Environment(
			processConstant.HostEnvironment,
			processConstant.PortEnvironment,
			processConstant.InsecureEnvironment,
		).String(),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return New(c, place)
}
