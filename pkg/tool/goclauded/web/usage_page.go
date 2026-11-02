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
	q *http.Request,
) {
	usage, _, e := s.usage.FindUsage()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderLivePage(
		w,
		constant.UsageTitle,
		constant.UsagePath,
		subscription.Query(constant.Usage),
		html.H3(gomponents.Text(constant.UsageTitle)),
		html.Div(
			extended.StreamSwap(constant.UsageChart),
			usage_page.Chart(s.usage.UsageWindow()),
		),
		usage_page.Note(usage, s.usage.UtilizationCredential()),
	)
}
