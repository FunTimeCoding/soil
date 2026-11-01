package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/message_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) messagesPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	page, e := strconv.Atoi(q.URL.Query().Get(constant.Page))

	if e != nil || page < 1 {
		page = 1
	}

	messages := s.service.AllMessages(
		constant.PageSize+1,
		(page-1)*constant.PageSize,
	)
	more := len(messages) > constant.PageSize

	if more {
		messages = messages[:constant.PageSize]
	}

	s.view.RenderPage(
		w,
		constant.MessagesTitle,
		constant.MessagesPath,
		html.H3(gomponents.Text(constant.MessagesTitle)),
		message_page.List(messages, page, more),
	)
}
