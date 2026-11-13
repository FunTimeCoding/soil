package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"net/http"
)

func (s *Server) memoryDetailPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	v, okay := s.visibleMemory(w, q)

	if !okay {
		return
	}

	statistic, summary, g := s.service.MemoryTokenStatistic(
		v.Memory.Identifier,
		v.Memory.Scope,
	)

	if g != nil {
		s.reporter.CaptureException(g)
	}

	p := memory_page.New(v)
	s.view.RenderSummaryPage(
		w,
		v.Memory.Name,
		constant.MemoriesPath,
		p.TokenStrip(statistic, summary),
		p.Detail(),
	)
}
