package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) issueCertificate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var node []gomponents.Node
	node = append(
		node,
		html.H1(gomponents.Text(constant.IssueCertificateTitle)),
	)

	if message := form.ErrorText(r); message != "" {
		node = append(node, layout.Alert(message))
	}

	node = append(node, issueCertificateForm(r))
	s.view.RenderPage(
		w,
		constant.IssueCertificateTitle,
		constant.IssueCertificatePath,
		node...,
	)
}
