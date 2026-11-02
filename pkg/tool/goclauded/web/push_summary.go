package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/usage_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (s *Server) pushSummary(w http.ResponseWriter) {
	usage, _, e := s.usage.FindUsage()

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	if items := usage_page.Summary(usage); len(items) > 0 {
		layout.PushEvent(
			w,
			webConstant.LayoutSummaryStrip,
			layout.SummaryStripContent(items),
		)
	}
}
