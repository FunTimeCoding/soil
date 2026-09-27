package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func requestSection(requests []*merge_request.Request) gomponents.Node {
	if len(requests) == 0 {
		return gomponents.Group(nil)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.RequestTitle)),
			requestTable(requests),
		},
	)
}
