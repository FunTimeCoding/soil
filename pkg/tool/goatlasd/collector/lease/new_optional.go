package lease

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewOptional() *Collector {
	token := environment.Optional(constant.TokenEnvironment)

	if token == "" {
		return nil
	}

	c, e := client.NewClientWithResponses(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		).String(),
		client.WithHTTPClient(web.StallClient()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return New(c)
}
