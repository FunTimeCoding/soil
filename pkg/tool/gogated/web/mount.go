package web

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.WithSession(s.requireAdmin)
	g.Session(route.Get(webConstant.PalettePath), s.registry.Serve())
	a := s.service.Authorization()
	g.Open(route.Get(webConstant.SignInPath), a.SignIn)
	g.Open(route.Get(webConstant.CallbackPath), a.Callback)
	g.Open(route.Get(webConstant.SignOutPath), a.SignOut)
	g.Session(route.Get(webConstant.RootPattern), s.clients)
	g.Session(route.Get(constant.CreatePath), s.create)
	g.Session(route.Post(constant.CreatePath), s.createSubmit)
	g.Session(
		route.Get(constant.ClientsPrefix, constant.IdentifierPattern),
		s.clientDetail,
	)
	g.Session(
		route.Post(
			constant.ClientsPrefix,
			constant.IdentifierPattern,
			constant.DeleteSuffix,
		),
		s.deleteClient,
	)
	g.Session(route.Get(constant.SessionsPath), s.sessions)
	g.Session(
		route.Post(constant.SessionsRevokeOthersPath),
		s.revokeOtherSessions,
	)
	g.Session(
		route.Post(
			constant.SessionsPath,
			stringsConstant.Slash,
			constant.IdentifierPattern,
			constant.DeleteSuffix,
		),
		s.deleteSession,
	)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
