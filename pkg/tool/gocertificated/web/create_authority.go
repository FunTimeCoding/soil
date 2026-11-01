package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web/authority_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) createAuthority(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.CreateAuthorityTitle,
		constant.CreateAuthorityPath,
		html.H1(gomponents.Text(constant.CreateAuthorityTitle)),
		authority_page.CreateForm(r.URL.Query(), form.ErrorText(r)),
	)
}
