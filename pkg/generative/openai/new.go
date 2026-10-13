package openai

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func New(token string) *Client {
	p := stall_client.New().Build()

	return &Client{
		client: openai.NewClient(
			option.WithAPIKey(token),
			option.WithHTTPClient(p.Client()),
		),
		pool: p,
	}
}
