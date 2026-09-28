package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store/fosite_client"
	"github.com/ory/fosite"
)

func (s *Store) GetClient(
	_ context.Context,
	identifier string,
) (fosite.Client, error) {
	var row client.Client
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if e != nil {
		return nil, fosite.ErrNotFound
	}

	return fosite_client.New(&row), nil
}
