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
	c := web.Client()

	if untrusted {
		c = web.InsecureClient()
	}

	options := []client.ClientOption{
		client.WithHTTPClient(c),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	}
	generated, e := client.NewClientWithResponses(base, options...)
	errors.PanicOnError(e)

	return &Client{
		generated: generated,
		base:      base,
		token:     token,
		untrusted: untrusted,
	}
}
