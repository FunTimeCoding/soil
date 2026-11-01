package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"io"
	"net/http"
)

func (c *Client) Listen(x context.Context) (io.ReadCloser, error) {
	q, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		join.Empty(c.base, constant.StreamPath),
		nil,
	)

	if e != nil {
		return nil, e
	}

	web.Bearer(q, c.token)
	r, e := c.pool.Client().Do(q)

	if e != nil {
		return nil, e
	}

	if r.StatusCode != http.StatusOK {
		defer errors.LogClose(r.Body)
		b, f := io.ReadAll(r.Body)

		if f != nil {
			return nil, f
		}

		return nil, refusal.New(constant.StreamOperation, r, b)
	}

	return r.Body, nil
}
