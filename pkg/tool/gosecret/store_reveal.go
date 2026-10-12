package gosecret

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
)

func (s *interfaceStore) Reveal(
	path string,
	key string,
) (string, error) {
	r, e := s.client.RevealSecretWithResponse(
		context.Background(),
		&client.RevealSecretParams{Path: path, Key: key},
	)

	if e != nil {
		return "", fmt.Errorf("reveal secret: %w", e)
	}

	if r.JSON200 == nil {
		return "", unexpected.Format("reveal secret %s: %s", r.Status(), r.Body)
	}

	return r.JSON200.Value, nil
}
