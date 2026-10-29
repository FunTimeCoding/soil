package service

import (
	"github.com/funtimecoding/soil/pkg/iterm/history"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
)

func (s *Service) ReadHistory(
	identifier string,
	count int,
) (*history.History, error) {
	if count <= 0 {
		count = constant.HistoryCount
	}

	return s.source.ReadHistory(identifier, count)
}
