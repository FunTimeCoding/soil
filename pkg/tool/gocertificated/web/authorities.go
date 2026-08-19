package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) authorities(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.store.Authorities()
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		constant.AuthoritiesTitle,
		constant.AuthoritiesPath,
		html.H1(gomponents.Text(constant.AuthoritiesTitle)),
		html.P(
			html.A(
				html.Href(constant.CreateAuthorityPath),
				gomponents.Text(constant.CreateAuthorityTitle),
			),
		),
		authoritiesTable(result),
	)
}
