package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
)

func (s *Service) Log(limit int) ([]*log_entry.Entry, error) {
	if limit == 0 {
		limit = constant.DefaultLogLimit
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*log_entry.Entry, error) {
			return s.client.Log(limit)
		},
	)
}
