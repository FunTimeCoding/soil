package openbao

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/openbao/openbao/api/v2"
)

func (c *Client) Read(path string) (map[string]string, error) {
	x := context.Background()

	if e := c.authenticate(x); e != nil {
		return nil, e
	}

	s, e := c.client.KVv2(constant.Mount).Get(x, path)

	if errors.Is(e, api.ErrSecretNotFound) {
		return nil, not_found.New(constant.SecretKind, path)
	}

	if e != nil {
		return nil, fmt.Errorf("read secret: %w", e)
	}

	result := map[string]string{}

	for k, v := range s.Data {
		result[k] = fmt.Sprint(v)
	}

	return result, nil
}
