package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
)

func location(s *symbol.Symbol) *server.Location {
	if s == nil {
		return nil
	}

	return &server.Location{Key: s.Key(), File: s.File, Line: s.Line}
}
