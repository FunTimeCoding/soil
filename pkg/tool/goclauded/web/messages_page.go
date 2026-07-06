package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) messagesPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	limit := 20
	skip := (page - 1) * limit
	messages := s.service.AllMessages(limit+1, skip)
	hasMore := len(messages) > limit

	if hasMore {
		messages = messages[:limit]
	}

	var rows []gomponents.Node

	for i := range messages {
		rows = append(rows, messageRow(&messages[i]))
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.MessagesTitle)),
	)

	if len(rows) == 0 {
		content = append(content, html.P(gomponents.Text("No messages.")))
	} else {
		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("From")),
						html.Th(gomponents.Text("To")),
						html.Th(gomponents.Text("Message")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	var navigation []gomponents.Node

	if page > 1 {
		navigation = append(
			navigation,
			html.A(
				gomponents.Attr("href", messagePaginationLink(page-1)),
				gomponents.Text("← Newer"),
			),
		)
	}

	if hasMore {
		if len(navigation) > 0 {
			navigation = append(navigation, gomponents.Text(" · "))
		}

		navigation = append(
			navigation,
			html.A(
				gomponents.Attr("href", messagePaginationLink(page+1)),
				gomponents.Text("Older →"),
			),
		)
	}

	if len(navigation) > 0 {
		content = append(content, html.P(gomponents.Group(navigation)))
	}

	s.view.RenderPage(
		w,
		constant.MessagesTitle,
		constant.MessagesPath,
		content...,
	)
}
