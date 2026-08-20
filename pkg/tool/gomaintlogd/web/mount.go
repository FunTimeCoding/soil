package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(
		route.Get(webConstant.PalettePath),
		palette.NewServe(s.registry),
	)
	m.HandleFunc(route.Get(webConstant.RootPattern), s.dashboard)
	m.HandleFunc(route.Get(constant.EntriesPath), s.entries)
	m.HandleFunc(
		route.Get(
			fmt.Sprintf("%s/{%s}", constant.EntryPath, constant.Identifier),
		),
		s.entryPage,
	)
	m.HandleFunc(route.Get(constant.AddEntryPath), s.add)
	m.HandleFunc(route.Post(constant.AddEntryPath), s.addSubmit)
	m.HandleFunc(route.Get(constant.DetailPath), s.detail)
	m.HandleFunc(route.Get(constant.EditPath), s.edit)
	m.HandleFunc(route.Post(constant.EditPath), s.editSubmit)
	m.HandleFunc(route.Post(constant.DeletePath), s.delete)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
