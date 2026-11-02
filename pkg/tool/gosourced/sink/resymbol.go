package sink

import (
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/system"
)

func (s *Sink) Resymbol(
	directory string,
	name string,
	newDirectory string,
	newName string,
) {
	if !system.InsideDirectory(s.root, directory) ||
		!system.InsideDirectory(s.root, newDirectory) {
		return
	}

	s.symbols = append(
		s.symbols,
		symbol_move.New(directory, name, newDirectory, newName),
	)
}
