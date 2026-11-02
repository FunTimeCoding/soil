package sink

import "github.com/funtimecoding/soil/pkg/lint/types/symbol_move"

func (s *Sink) Symbols() []*symbol_move.Move {
	return s.symbols
}
