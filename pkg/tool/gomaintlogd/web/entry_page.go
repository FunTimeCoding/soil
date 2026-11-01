package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	"net/http"
)

func (s *Server) entryPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := strings.ToUnsignedInteger(r.PathValue(constant.Identifier), 0)

	if identifier == 0 {
		s.entryMissing(w)

		return
	}

	e, f := s.service.Entry(identifier)

	if not_found.Is(f) {
		s.entryMissing(w)

		return
	}

	errors.PanicOnError(f)
	s.view.RenderPage(
		w,
		constant.EntryTitle,
		entry_page.Locator(identifier),
		entry_page.Permalink(e),
	)
}
