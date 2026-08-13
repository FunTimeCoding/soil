package openai

import (
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func NewBase(base string) *Client {
	return &Client{
		client: openai.NewClient(
			option.WithBaseURL(base),
			option.WithAPIKey("none"),
		),
	}
}
