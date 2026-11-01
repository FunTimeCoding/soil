package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/usage_page"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/subscription"
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
		subscription.Query(constant.Usage),
		html.H3(gomponents.Text(constant.UsageTitle)),
		html.Div(
			extended.StreamSwap(constant.UsageChart),
			usage_page.Chart(s.service.UsageWindow()),
		),
		usage_page.Note(s.service.Usage(), s.service.UtilizationCredential()),
	)
}
