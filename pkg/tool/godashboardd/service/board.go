package service

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

func (s *Service) Board() *board.Board {
	return s.board
}
