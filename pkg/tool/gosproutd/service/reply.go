package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/turn"
)

func (s *Service) Reply(
	identifier uint,
	content string,
) (*turn.Turn, error) {
	return s.AddTurn(identifier, constant.AuthorUser, content)
}
