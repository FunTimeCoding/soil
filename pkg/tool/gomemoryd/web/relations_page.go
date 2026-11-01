package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/relation_page"
	"net/http"
)

func (s *Server) relationsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	relations, e := s.service.VisibleRelations()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.RelationsTitle,
		constant.RelationsPath,
		relation_page.List(relations, q.URL.Query().Get(constant.Type)),
	)
}
