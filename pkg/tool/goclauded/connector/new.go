package connector

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func New(
	base string,
	untrusted bool,
	token string,
) *Client {
	options := []client.ClientOption{
		client.WithRequestEditorFn(web.BearerEditor(token)),
	}

	if untrusted {
		options = append(options, client.WithHTTPClient(web.InsecureClient()))
	}

	generated, e := client.NewClientWithResponses(base, options...)
	errors.PanicOnError(e)

	return &Client{generated: generated}
}
