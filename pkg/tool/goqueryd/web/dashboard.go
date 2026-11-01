package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/collection_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	q *http.Request,
) {
	status, e := s.service.Status()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPageWithSummary(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		[]string{
			fmt.Sprintf("%d documents", status.TotalDocuments),
			fmt.Sprintf("%d embeddings", status.TotalEmbeddings),
			fmt.Sprintf("%d pending", status.PendingEmbeddings),
			fmt.Sprintf("%d failed", status.FailedEmbeddings),
		},
		html.H3(gomponents.Text(constant.CollectionsTitle)),
		collection_page.List(status.Collections),
	)
}
