package process

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewOptional() *Collector {
	place := environment.Optional(constant.PlaceEnvironment)
	token := environment.Optional(goprocessdConstant.TokenEnvironment)

	if place == "" || token == "" {
		return nil
	}

	c, e := client.NewClientWithResponses(
		locator.Environment(
			goprocessdConstant.HostEnvironment,
			goprocessdConstant.PortEnvironment,
			goprocessdConstant.InsecureEnvironment,
		).String(),
		client.WithHTTPClient(web.StallClient()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return New(c, place)
}
