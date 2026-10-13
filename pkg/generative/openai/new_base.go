package openai

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func NewBase(base string) *Client {
	p := stall_client.New().Long().Build()

	return &Client{
		client: openai.NewClient(
			option.WithBaseURL(base),
			option.WithAPIKey("none"),
			option.WithHTTPClient(p.Client()),
		),
		pool: p,
	}
}
