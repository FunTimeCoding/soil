package raid_parser

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	insecure bool,
	token string,
) *Client {
	l := locator.New(host)

	if insecure {
		l.Insecure()
	}

	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
