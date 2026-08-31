package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) usagePage(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderLivePage(
		w,
		constant.UsageTitle,
		constant.UsagePath,
		join.Empty("subscribe=", constant.Usage),
		html.H3(gomponents.Text(constant.UsageTitle)),
		html.Div(
			gomponents.Attr("sse-swap", constant.UsageChart),
			s.usageChart(),
		),
	)
}
