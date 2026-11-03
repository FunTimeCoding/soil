package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) Play(
	_ context.Context,
	identifier string,
	b client.PlayJSONRequestBody,
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

	if len(b.ItemIds) == 0 {
		return validation.New("item_ids is required")
	}

	c.Played[identifier] = append(c.Played[identifier], b)

	return nil
}
