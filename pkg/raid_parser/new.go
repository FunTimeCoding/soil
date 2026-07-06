package raid_parser

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	insecure bool,
) *Client {
	l := locator.New(host)

	if insecure {
		l.Insecure()
	}

	c, e := client.NewClientWithResponses(l.String())
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
