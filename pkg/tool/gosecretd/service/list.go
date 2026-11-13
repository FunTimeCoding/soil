package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"strings"
)

func (s *Service) List(path string) ([]string, error) {
	path = strings.TrimSuffix(path, constant.Folder)

	if path != "" {
		if e := validatePath(path); e != nil {
			return nil, e
		}
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]string, error) {
			return s.client.List(path)
		},
	)
}
