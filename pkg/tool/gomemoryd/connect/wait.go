package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
	memoryClient "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/peer"
	"net/http"
)

func Wait(
	l *logger.Logger,
	h *http.Client,
) memoryClient.Client {
	base := locator.Environment(
		constant.HostEnvironment,
		constant.PortEnvironment,
		constant.InsecureEnvironment,
	).String()
	c, e := generated.NewClient(
		base,
		generated.WithHTTPClient(h),
		generated.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	errors.PanicOnError(
		peer.Await(
			l,
			constant.Identity.Name(),
			base,
			func() error {
				r, f := c.GetVersions(
					context.Background(),
					&generated.GetVersionsParams{Since: "2099-01-01T00:00:00Z"},
				)

				if f != nil {
					return f
				}

				errors.PanicClose(r.Body)

				return nil
			},
		),
	)

	return memoryClient.NewRestClient(c)
}
