package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	openbaoConstant "github.com/funtimecoding/soil/pkg/openbao/constant"
)

func (c *Client) secret(path string) (map[string]string, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.Secrets[path]

	if !okay {
		return nil, not_found.New(openbaoConstant.SecretKind, path)
	}

	return result, nil
}
