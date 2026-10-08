package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strings"
)

func (s *Server) hits(
	w http.ResponseWriter,
	r *http.Request,
) {
	session := r.PathValue(constant.Identifier)
	query := r.URL.Query().Get(constant.QueryField)
	kinds := r.URL.Query()[constant.KindField]
	around := r.URL.Query().Get(constant.Around)
	blocks, e := s.service.ConversationBlocks(session, kinds)
	errors.PanicOnError(e)
	matching, f := s.service.MatchingBlocks(session, query, kinds)
	errors.PanicOnError(f)
	terms := strings.Fields(query)
	var nodes []gomponents.Node

	for _, b := range blocks {
		nodes = append(
			nodes,
			hitBlock(b, terms, matching[b.Identifier], b.Identifier == around),
		)
	}

	nodes = append(
		nodes,
		html.Div(html.ID("hit-counter"), html.Class("hit-counter")),
	)
	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(gomponents.Group(nodes).Render(w))
}
