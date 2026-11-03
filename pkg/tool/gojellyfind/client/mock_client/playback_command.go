package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) PlaybackCommand(
	_ context.Context,
	identifier string,
	b client.PlaybackCommandJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return e
	}

	if identifier == "" {
		return not_found.New("session", identifier)
	}

	if b.Command == "" {
		return validation.New("command is required")
	}

	c.Commanded[identifier] = append(c.Commanded[identifier], b)

	return nil
}
