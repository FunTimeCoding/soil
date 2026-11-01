package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"net/http"
	"strconv"
)

func (s *Server) memoryDetailPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	raw := q.PathValue(constant.Identifier)
	identifier, e := strconv.ParseInt(raw, 10, 64)

	if e != nil {
		s.view.Answer(w, q, s.reporter, not_found.New(constant.MemoryKind, raw))

		return
	}

	v, f := s.service.MemoryView(identifier)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	statistic, summary, g := s.service.MemoryTokenStatistic(
		identifier,
		v.Memory.Scope,
	)

	if g != nil {
		s.reporter.CaptureException(g)
	}

	p := memory_page.New(v)
	s.view.RenderPageWithSummary(
		w,
		v.Memory.Name,
		constant.MemoriesPath,
		p.TokenStrip(statistic, summary),
		p.Detail(),
	)
}
