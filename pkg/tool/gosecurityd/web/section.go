package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/flag_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func (s *Server) section(
	values url.Values,
	alert string,
	show string,
) (gomponents.Node, error) {
	o, e := s.overview.Overview(show)

	if e != nil {
		return nil, e
	}

	return html.Div(
		html.ID(constant.FlagsMark),
		flag_page.Tiles(o),
		flag_page.FlagForm(values, alert),
		flag_page.Filter(show),
		flag_page.Table(o.Rows),
	), nil
}
