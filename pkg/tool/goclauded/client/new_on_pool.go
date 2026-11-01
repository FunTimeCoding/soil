package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func newOnPool(
	b *stall_client.Builder,
	base string,
	untrusted bool,
	token string,
) *Client {
	if untrusted {
		b.Insecure()
	}

	p := b.Build()
	options := []client.ClientOption{
		client.WithHTTPClient(p.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	}
	generated, e := client.NewClientWithResponses(base, options...)
	errors.PanicOnError(e)

	return &Client{generated: generated, base: base, token: token, pool: p}
}
