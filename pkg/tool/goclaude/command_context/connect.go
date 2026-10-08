package command_context

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func connect(
	address string,
	requester *http.Client,
	token string,
) *client.ClientWithResponses {
	result, e := client.NewClientWithResponses(
		address,
		client.WithHTTPClient(requester),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return result
}
