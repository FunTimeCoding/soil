package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
)

func (s *Service) Reveal(
	path string,
	key string,
) (string, error) {
	if e := validatePath(path); e != nil {
		return "", e
	}

	if key == "" {
		return "", validation.New("key is required")
	}

	values, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (map[string]string, error) {
			return s.client.Read(path)
		},
	)

	if e != nil {
		return "", e
	}

	result, okay := values[key]

	if !okay {
		return "", not_found.New(constant.KeyKind, key)
	}

	return result, nil
}
