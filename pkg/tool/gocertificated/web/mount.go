package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(route.Get(webConstant.SignInPath), s.signIn)
	m.HandleFunc(route.Get(webConstant.CallbackPath), s.callback)
	m.HandleFunc(route.Get(webConstant.SignOutPath), s.signOut)
	m.HandleFunc(
		route.Get(webConstant.PalettePath),
		s.require(palette.NewServe(s.registry)),
	)
	m.HandleFunc(route.Get(webConstant.RootPattern), s.require(s.dashboard))
	m.HandleFunc(route.Get(constant.AuthoritiesPath), s.require(s.authorities))
	m.HandleFunc(
		route.Get(constant.CertificatesPath),
		s.require(s.certificates),
	)
	m.HandleFunc(
		route.Get(constant.CreateAuthorityPath),
		s.require(s.createAuthority),
	)
	m.HandleFunc(
		route.Post(constant.CreateAuthorityPath),
		s.require(s.createAuthoritySubmit),
	)
	m.HandleFunc(
		route.Get(constant.IssueCertificatePath),
		s.require(s.issueCertificate),
	)
	m.HandleFunc(
		route.Post(constant.IssueCertificatePath),
		s.require(s.issueCertificateSubmit),
	)
	m.HandleFunc(route.Post(constant.PublishPath), s.require(s.publishSubmit))
	m.HandleFunc(route.Get(constant.RootPath), s.root)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
