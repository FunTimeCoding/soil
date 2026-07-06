package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"maragu.dev/gomponents"
)

func (s *Server) recentTable() gomponents.Node {
	f := store.NewFilter()
	f.Limit = 50
	entries, e := s.store.List(f)
	errors.PanicOnError(e)

	return entriesTable(entries)
}
