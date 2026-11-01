package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	"net/http"
)

func (s *Server) entryMissing(w http.ResponseWriter) {
	s.view.RenderPageWithStatus(
		w,
		http.StatusNotFound,
		constant.EntryTitle,
		constant.EntriesPath,
		entry_page.Missing(),
	)
}
