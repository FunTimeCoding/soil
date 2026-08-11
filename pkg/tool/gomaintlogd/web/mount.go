package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc(fmt.Sprintf("GET %s", constant.EntriesPath), s.entries)
	m.HandleFunc(
		fmt.Sprintf("GET %s/{%s}", constant.EntryPath, constant.Identifier),
		s.entryPage,
	)
	m.HandleFunc(fmt.Sprintf("GET %s", constant.AddEntryPath), s.add)
	m.HandleFunc(fmt.Sprintf("POST %s", constant.AddEntryPath), s.addSubmit)
	m.HandleFunc(fmt.Sprintf("GET %s", constant.DetailPath), s.detail)
	m.HandleFunc(fmt.Sprintf("GET %s", constant.EditPath), s.edit)
	m.HandleFunc(fmt.Sprintf("POST %s", constant.EditPath), s.editSubmit)
	m.HandleFunc(fmt.Sprintf("POST %s", constant.DeletePath), s.delete)
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
