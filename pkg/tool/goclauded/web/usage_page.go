package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
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
		html.Div(extended.StreamSwap(constant.UsageChart), s.usageChart()),
		s.usageNote(),
	)
}
