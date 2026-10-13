package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) requestSection(requests []*request_entry.Entry) gomponents.Node {
	if len(requests) == 0 {
		return gomponents.Group(nil)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.RequestTitle)),
			requestTable(requests, s.multiple()),
		},
	)
}
