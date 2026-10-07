package virtual_file_system

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system/pending_move"

func (s *System) Move(
	from string,
	to string,
) {
	s.moves = append(s.moves, pending_move.Move{From: from, To: to})
}
