package openbao

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
)

func (c *Client) Put(
	path string,
	values map[string]string,
) error {
	x := context.Background()

	if e := c.authenticate(x); e != nil {
		return e
	}

	payload := map[string]any{}

	for k, v := range values {
		payload[k] = v
	}

	if _, e := c.client.KVv2(constant.Mount).Put(x, path, payload); e != nil {
		return fmt.Errorf("put secret: %w", e)
	}

	return nil
}
