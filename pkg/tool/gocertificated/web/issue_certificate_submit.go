package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) issueCertificateSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	b := &server.CertificateBody{
		Authority:  r.PostFormValue(constant.AuthorityParameter),
		Kind:       server.LeafKind(r.PostFormValue(constant.KindParameter)),
		CommonName: r.PostFormValue(constant.CommonNameParameter),
	}
	optionalList(&b.Host, r.PostFormValue(constant.HostParameter))
	result, key, e := s.service.IssueCertificate(b)

	if e != nil {
		form.Redirect(
			w,
			r,
			constant.IssueCertificatePath,
			e.Error(),
			r.PostForm,
		)

		return
	}

	s.view.RenderPage(
		w,
		constant.IssueCertificateTitle,
		constant.IssueCertificatePath,
		html.H1(gomponents.Text(constant.IssueCertificateTitle)),
		html.P(
			html.Strong(
				gomponents.Text(
					"This key is shown once and is not stored. Copy it now.",
				),
			),
		),
		html.H2(gomponents.Text("Certificate")),
		html.Pre(html.Code(gomponents.Text(result.Certificate))),
		html.H2(gomponents.Text("Private key")),
		html.Pre(html.Code(gomponents.Text(key))),
	)
}
