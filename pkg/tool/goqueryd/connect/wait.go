package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/peer"
	"net/http"
)

func Wait(
	l *logger.Logger,
	h *http.Client,
) *client.Client {
	base := locator.Environment(
		constant.HostEnvironment,
		constant.PortEnvironment,
		constant.InsecureEnvironment,
	).String()
	c, e := client.NewClient(
		base,
		client.WithHTTPClient(h),
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	errors.PanicOnError(
		peer.Await(
			l,
			queryd.Identity.Name(),
			base,
			func() error {
				r, f := c.GetStatus(context.Background())

				if f != nil {
					return f
				}

				errors.PanicClose(r.Body)

				return nil
			},
		),
	)

	return c
}
