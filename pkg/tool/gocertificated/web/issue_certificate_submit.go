package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web/certificate_page"
	"github.com/funtimecoding/soil/pkg/web/form"
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

	if e != nil && classify.Reportable(e) {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

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
		certificate_page.Issued(result.Certificate, key),
	)
}
