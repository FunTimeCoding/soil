package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.WithSession(s.requireAdmin)
	g.Session(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(webConstant.SignInPath), s.signIn)
	g.Open(route.Get(webConstant.CallbackPath), s.callback)
	g.Open(route.Get(webConstant.SignOutPath), s.signOut)
	g.Session(route.Get(webConstant.RootPattern), s.clients)
	g.Session(route.Get(constant.CreatePath), s.create)
	g.Session(route.Post(constant.CreatePath), s.createSubmit)
	g.Session(route.Get("/clients/{identifier}"), s.clientDetail)
	g.Session(route.Post("/clients/{identifier}/delete"), s.deleteClient)
	g.Session(route.Get(constant.SessionsPath), s.sessions)
	g.Session(
		route.Post(constant.SessionsRevokeOthersPath),
		s.revokeOtherSessions,
	)
	g.Session(
		route.Post(constant.SessionsPath, "/{identifier}/delete"),
		s.deleteSession,
	)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
