package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/flag_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) affectedCell(name string) gomponents.Node {
	a, e := s.system.Affected(name)

	if e != nil {
		return html.P(html.Small(gomponents.Text(e.Error())))
	}

	return flag_page.Affected(a)
}
