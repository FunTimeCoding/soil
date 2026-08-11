package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
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

	e, f := s.store.Get(identifier)

	if not_found.Is(f) {
		s.entryMissing(w)

		return
	}

	errors.PanicOnError(f)
	s.view.RenderPage(
		w,
		constant.EntryTitle,
		entryLocator(identifier),
		html.H1(gomponents.Textf("%s %d", constant.EntryTitle, e.ID)),
		html.Div(html.Class("detail-content"), entryFields(e)),
		html.A(
			gomponents.Attr("role", "button"),
			html.Class("outline secondary"),
			html.Href(constant.EntriesPath),
			gomponents.Text(constant.EntriesTitle),
		),
	)
}
