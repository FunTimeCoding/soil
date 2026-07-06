package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/peek"

func (c *Client) Peek(sessionIdentifier string) *peek.Peek {
	if messages, okay := c.UserMessages[sessionIdentifier]; okay {
		result := peek.New()

		for _, m := range messages {
			result.Entries = append(
				result.Entries,
				peek.Entry{UserText: m},
			)
			result.UserMessageCount++
		}

		return result
	}

	return peek.Stub()
}
