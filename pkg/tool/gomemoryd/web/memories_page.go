package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"net/http"
	"strconv"
)

func (s *Server) memoriesPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	values := q.URL.Query()
	page, e := strconv.Atoi(values.Get(constant.Page))

	if e != nil {
		page = 1
	}

	l, f := s.service.MemoryList(
		values.Get(constant.Type),
		values.Get(constant.Tag),
		values.Get(constant.Scope),
		page,
	)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	s.view.RenderPage(
		w,
		constant.MemoriesTitle,
		constant.MemoriesPath,
		memory_page.List(l),
	)
}
