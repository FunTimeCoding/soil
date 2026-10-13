package gosecret

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func interfaceClient(
	t *terminal.Terminal,
	h *http.Client,
) *client.ClientWithResponses {
	result, e := client.NewClientWithResponses(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		).String(),
		client.WithHTTPClient(h),
		client.WithRequestEditorFn(
			web.BearerEditor(t.Required(constant.TokenEnvironment)),
		),
	)
	errors.PanicOnError(e)

	return result
}
