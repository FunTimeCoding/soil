package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"net/http"
)

func (s *Server) entryFromQuery(r *http.Request) (*entry.Entry, error) {
	identifier, e := identifierFromQuery(r)

	if e != nil {
		return nil, e
	}

	return s.service.Entry(identifier)
}
