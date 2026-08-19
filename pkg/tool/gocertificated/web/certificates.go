package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) certificates(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.store.Certificates(store.NewFilter())
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		constant.CertificatesTitle,
		constant.CertificatesPath,
		html.H1(gomponents.Text(constant.CertificatesTitle)),
		html.P(
			html.A(
				html.Href(constant.IssueCertificatePath),
				gomponents.Text(constant.IssueCertificateTitle),
			),
		),
		certificatesTable(result),
	)
}
