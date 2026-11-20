package web

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/flag_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
)

func (s *Server) lookup(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get(generativeConstant.ParameterIdentifier)
	nodes := []gomponents.Node{
		html.H1(gomponents.Text(constant.LookupTitle)),
		flag_page.LookupForm(name),
	}

	if name != "" {
		v, e := s.vulnerability.Vulnerability(r.Context(), name)
		nodes = append(
			nodes,
			flag_page.VulnerabilityCard(name, v, e),
			html.H2(gomponents.Text("Debian")),
			s.debianCell(name),
			html.Div(
				html.ID(constant.FlagsMark),
				flag_page.FlagForm(
					url.Values{generativeConstant.ParameterIdentifier: {name}},
					"",
				),
			),
		)
	}

	s.view.RenderPage(w, constant.LookupTitle, constant.LookupPath, nodes...)
}
