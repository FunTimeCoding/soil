package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"net/http"
	"net/url"
)

func (s *Server) flagSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	name := r.FormValue(generativeConstant.ParameterIdentifier)
	_, e := s.flag.Flag(r.Context(), name, r.FormValue(constant.NoteField))

	if refused(e) {
		s.renderSection(w, r, r.PostForm, e.Error(), "")

		return
	}

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.renderSection(w, r, url.Values{}, "", fmt.Sprintf("Flagged %s", name))
}
