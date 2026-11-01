package certificate_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Issued(
	certificate string,
	key string,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H1(gomponents.Text(constant.IssueCertificateTitle)),
			html.P(
				html.Strong(
					gomponents.Text(
						"This key is shown once and is not stored. Copy it now.",
					),
				),
			),
			html.H2(gomponents.Text("Certificate")),
			html.Pre(html.Code(gomponents.Text(certificate))),
			html.H2(gomponents.Text("Private key")),
			html.Pre(html.Code(gomponents.Text(key))),
		},
	)
}
