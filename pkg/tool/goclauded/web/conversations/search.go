package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strings"
)

func (s *Server) search(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query().Get(constant.QueryField)
	kinds := r.URL.Query()[constant.KindField]
	w.Header().Set(web.ContentType, "text/html; charset=utf-8")

	if strings.TrimSpace(query) == "" {
		errors.PanicOnError(gomponents.Group(s.sidebarNodes(0)).Render(w))

		return
	}

	found, e := s.service.SearchConversations(query, kinds, 0)
	errors.PanicOnError(e)
	var nodes []gomponents.Node

	if indexed, total := s.service.SearchProgress(); indexed < total {
		nodes = append(
			nodes,
			html.Small(
				html.Class("search-progress"),
				gomponents.Text(
					fmt.Sprintf("indexing %d/%d conversations", indexed, total),
				),
			),
		)
	}

	if len(found) == 0 {
		nodes = append(
			nodes,
			html.Small(
				html.Class("search-progress"),
				gomponents.Text("No conversation holds every term."),
			),
		)
	}

	for i, c := range found {
		nodes = append(nodes, searchResult(c, query, kinds, i == 0))
	}

	errors.PanicOnError(gomponents.Group(nodes).Render(w))
}
