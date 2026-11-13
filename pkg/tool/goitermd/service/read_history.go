package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/iterm/history"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
)

func (s *Service) ReadHistory(
	identifier string,
	count int,
) (*history.History, error) {
	if count <= 0 {
		count = constant.HistoryCount
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*history.History, error) {
			return s.client.ReadHistory(identifier, count)
		},
	)
}
