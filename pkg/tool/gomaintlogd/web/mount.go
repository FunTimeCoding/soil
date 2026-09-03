package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.Open(route.Get(constant.EntriesPath), s.entries)
	g.Open(
		route.Get(
			fmt.Sprintf("%s/{%s}", constant.EntryPath, constant.Identifier),
		),
		s.entryPage,
	)
	g.Open(route.Get(constant.AddEntryPath), s.add)
	g.Open(route.Post(constant.AddEntryPath), s.addSubmit)
	g.Open(route.Get(constant.DetailPath), s.detail)
	g.Open(route.Get(constant.EditPath), s.edit)
	g.Open(route.Post(constant.EditPath), s.editSubmit)
	g.Open(route.Post(constant.DeletePath), s.delete)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
