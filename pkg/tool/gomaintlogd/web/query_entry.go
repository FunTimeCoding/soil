package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"net/http"
)

func (s *Server) queryEntry(r *http.Request) (*entry.Entry, error) {
	identifier, e := queryIdentifier(r)

	if e != nil {
		return nil, e
	}

	return s.service.Entry(identifier)
}
