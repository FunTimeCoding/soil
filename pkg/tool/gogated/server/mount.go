package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get("/.well-known/oauth-authorization-server"), s.discovery)
	g.Open(route.Get("/.well-known/openid-configuration"), s.discovery)
	g.Open(route.Post("/register"), s.register)
	g.Open(route.Get("/authorize"), s.authorizeGet)
	g.Open(route.Post("/authorize"), s.authorizePost)
	g.Open(route.Post("/token"), s.token)
	g.Open(route.Get("/jwks"), s.signingKeys)
	g.Open(route.Get(constant.LogoutPath), s.logoutGet)
	g.Open(route.Post(constant.LogoutPath), s.logoutPost)
}
