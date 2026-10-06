package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"strings"
)

func (c *Client) Stream(
	x context.Context,
	receive func([]client.Claim),
) error {
	q, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		join.Empty(c.base, constant.StreamPath),
		nil,
	)

	if e != nil {
		return fmt.Errorf("stream request: %w", e)
	}

	web.Bearer(q, c.token)
	r, e := web.Client().Do(q)

	if e != nil {
		return fmt.Errorf("stream: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("stream: %s", r.Status)
	}

	s := bufio.NewScanner(r.Body)

	for s.Scan() {
		payload, found := strings.CutPrefix(s.Text(), constant.EventPrefix)

		if !found {
			continue
		}

		var v []client.Claim

		if f := notation.Decode(payload, &v); f != nil {
			return fmt.Errorf("stream event: %w", f)
		}

		receive(v)
	}

	return s.Err()
}
