package web

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/history_page"
	"maragu.dev/gomponents"
)

func (s *Server) activity(
	x context.Context,
	kinds []string,
) (gomponents.Node, error) {
	entries, e := s.service.Timeline(
		x,
		argument.Timeline{Kinds: kinds, Limit: 10},
	)

	if e != nil {
		return nil, e
	}

	return history_page.Table(entries, false), nil
}
