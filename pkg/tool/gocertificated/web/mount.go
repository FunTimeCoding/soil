package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.WithSession(s.require)
	g.Open(route.Get(webConstant.SignInPath), s.signIn)
	g.Open(route.Get(webConstant.CallbackPath), s.callback)
	g.Open(route.Get(webConstant.SignOutPath), s.signOut)
	g.Session(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Session(route.Get(webConstant.RootPattern), s.dashboard)
	g.Session(route.Get(constant.AuthoritiesPath), s.authorities)
	g.Session(route.Get(constant.CertificatesPath), s.certificates)
	g.Session(route.Get(constant.CreateAuthorityPath), s.createAuthority)
	g.Session(route.Post(constant.CreateAuthorityPath), s.createAuthoritySubmit)
	g.Session(route.Get(constant.IssueCertificatePath), s.issueCertificate)
	g.Session(
		route.Post(constant.IssueCertificatePath),
		s.issueCertificateSubmit,
	)
	g.Session(route.Post(constant.PublishPath), s.publishSubmit)
	g.Open(route.Get(constant.RootPath), s.root)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
