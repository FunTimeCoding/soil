package anthropic

import (
	"context"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/message"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"log"
)

func (c *Client) StreamMessage(
	x context.Context,
	v []*message.Message,
	system string,
	model string,
	m constant.AnthropicMode,
) <-chan string {
	h := make(chan string, 64)
	go func() {
		defer close(h)
		s := c.client.Messages.NewStreaming(
			x,
			anthropic.MessageNewParams{
				Model:     model,
				MaxTokens: 8192,
				System:    buildSystem(system, m),
				Messages:  buildMessages(v, m),
			},
		)

		for s.Next() {
			e := s.Current()

			if e.Type == "content_block_delta" &&
				e.Delta.Type == "text_delta" {
				select {
				case h <- e.Delta.Text:
				case <-x.Done():
					return
				}
			}
		}

		if e := s.Err(); e != nil {
			log.Printf("stream error: %v", e)
		}

		errors.LogClose(s)
	}()

	return h
}
