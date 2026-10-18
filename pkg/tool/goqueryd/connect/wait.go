package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func Wait(
	l *logger.Logger,
	h *http.Client,
) *client.ClientWithResponses {
	base := locator.Environment(
		constant.HostEnvironment,
		constant.PortEnvironment,
		constant.InsecureEnvironment,
	).String()
	c, e := client.NewClientWithResponses(
		base,
		client.WithHTTPClient(h),
		client.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	errors.PanicOnError(
		peer.Await(
			context.Background(),
			l,
			queryd.Identity.Name(),
			base,
			func(x context.Context) error {
				r, f := c.GetStatusWithResponse(x)

				if f != nil {
					return f
				}

				if r.JSON200 == nil {
					return refusal.New("index status", r.HTTPResponse, r.Body)
				}

				return nil
			},
		),
	)

	return c
}
