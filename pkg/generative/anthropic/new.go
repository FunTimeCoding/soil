package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go/v2"
	"github.com/anthropics/anthropic-sdk-go/v2/option"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(token string) *Client {
	p := stall_client.New().Build()

	return &Client{
		client: anthropic.NewClient(
			option.WithAPIKey(token),
			option.WithHTTPClient(p.Client()),
		),
		pool: p,
	}
}
