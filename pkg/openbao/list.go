package openbao

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"slices"
)

func (c *Client) List(path string) ([]string, error) {
	if e := c.authenticate(context.Background()); e != nil {
		return nil, e
	}

	s, e := c.client.Logical().List(
		fmt.Sprintf(constant.MetadataFormat, constant.Mount, path),
	)

	if e != nil {
		return nil, fmt.Errorf("list secrets: %w", e)
	}

	if s == nil {
		return []string{}, nil
	}

	keys, _ := s.Data[constant.KeysField].([]any)
	result := []string{}

	for _, k := range keys {
		result = append(result, fmt.Sprint(k))
	}

	slices.Sort(result)

	return result, nil
}
