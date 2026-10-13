package hetzner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func New(
	token string,
	o ...Option,
) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	result.pool = stall_client.New().Build()
	result.client = hcloud.NewClient(
		hcloud.WithToken(token),
		hcloud.WithHTTPClient(result.pool.Client()),
	)

	return result
}
