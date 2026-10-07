package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/filter"
	"maragu.dev/gomponents"
)

func (s *Server) recentTable() gomponents.Node {
	f := filter.New()
	f.Limit = 50
	entries, e := s.store.List(f)
	errors.PanicOnError(e)

	return entriesTable(entries)
}
