package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
)

func (c *Client) RevealSecret(
	_ context.Context,
	p *client.RevealSecretParams,
) (*client.SecretValue, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	values, e := c.secret(p.Path)

	if e != nil {
		return nil, e
	}

	if p.Key == c.Tamper {
		return &client.SecretValue{Value: "tampered"}, nil
	}

	result, okay := values[p.Key]

	if !okay {
		return nil, not_found.New(constant.KeyKind, p.Key)
	}

	return &client.SecretValue{Value: result}, nil
}
