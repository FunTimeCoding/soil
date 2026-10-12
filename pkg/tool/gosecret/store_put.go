package gosecret

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
)

func (s *interfaceStore) Put(
	path string,
	values map[string]string,
) ([]string, error) {
	r, e := s.client.PutSecretWithResponse(
		context.Background(),
		&client.PutSecretParams{Path: path},
		client.PutSecretJSONRequestBody{Values: values},
	)

	if e != nil {
		return nil, fmt.Errorf("put secret: %w", e)
	}

	if r.JSON200 == nil {
		return nil, unexpected.Format("put secret %s: %s", r.Status(), r.Body)
	}

	return r.JSON200.Keys, nil
}
