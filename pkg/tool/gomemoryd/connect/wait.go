package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/system/environment"
	memoryClient "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
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
	c, e := generated.NewClientWithResponses(
		base,
		generated.WithHTTPClient(h),
		generated.WithRequestEditorFn(
			web.BearerEditor(environment.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)
	result := memoryClient.NewRestClient(c, base)
	errors.PanicOnError(
		peer.Await(
			context.Background(),
			l,
			constant.Identity.Name(),
			base,
			result.Reachable,
		),
	)

	return result
}
