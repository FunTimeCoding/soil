package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	base string,
	untrusted bool,
	token string,
	o ...Option,
) *Client {
	result := &Client{base: base, token: token}

	for _, f := range o {
		f(result)
	}

	b := stall_client.New().Traced().Timeout(result.timeout)

	if result.long {
		b.Long()
	}

	if untrusted {
		b.Insecure()
	}

	result.pool = b.Build()
	generated, e := client.NewClientWithResponses(
		base,
		client.WithHTTPClient(result.pool.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)
	result.generated = generated

	return result
}
