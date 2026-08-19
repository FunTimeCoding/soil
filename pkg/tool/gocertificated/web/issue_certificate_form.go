package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func issueCertificateForm(r *http.Request) gomponents.Node {
	return html.Form(
		html.Method(http.MethodPost),
		html.Action(constant.IssueCertificatePath),
		field(r, constant.AuthorityParameter, "Authority", "cluster or host"),
		field(r, constant.KindParameter, "Kind", "server or client"),
		field(r, constant.CommonNameParameter, "Common name", ""),
		field(
			r,
			constant.HostParameter,
			"Hosts",
			"comma separated names and addresses",
		),
		html.Button(html.Type("submit"), gomponents.Text("Issue")),
	)
}
