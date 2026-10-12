package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"maps"
)

func (c *Client) Read(path string) (map[string]string, error) {
	values, okay := c.secrets[path]

	if !okay {
		return nil, not_found.New(constant.SecretKind, path)
	}

	return maps.Clone(values), nil
}
