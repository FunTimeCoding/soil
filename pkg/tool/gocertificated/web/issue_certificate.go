package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web/certificate_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) issueCertificate(
	w http.ResponseWriter,
	r *http.Request,
) {
	authorities, e := s.service.Authorities()
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		constant.IssueCertificateTitle,
		constant.IssueCertificatePath,
		html.H1(gomponents.Text(constant.IssueCertificateTitle)),
		certificate_page.IssueForm(
			authorities,
			r.URL.Query(),
			form.ErrorText(r),
		),
	)
}
