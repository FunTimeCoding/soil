package graph

import (
	"context"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/microsoft/graph/response"
	"net/url"
)

func (c *Client) Reachable(_ context.Context) error {
	return c.Get(
		constant.ProfilePath,
		url.Values{constant.SelectFilter: {"id"}},
		response.NewProfile(),
	)
}
