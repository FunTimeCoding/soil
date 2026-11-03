package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) SetVolume(
	_ context.Context,
	identifier string,
	b client.SetVolumeJSONRequestBody,
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

	if b.Level < 0 || b.Level > 100 {
		return validation.New("level must be between 0 and 100")
	}

	c.Volumes[identifier] = append(c.Volumes[identifier], b.Level)

	return nil
}
