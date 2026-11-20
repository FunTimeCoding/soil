package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"net/http"
	"net/url"
)

func (s *Server) mitigate(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	name := r.FormValue(generativeConstant.ParameterIdentifier)
	_, e := s.flag.Mitigate(name, r.FormValue("mitigation"))

	if refused(e) {
		s.renderSection(w, r, url.Values{}, e.Error(), "")

		return
	}

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.renderSection(w, r, url.Values{}, "", fmt.Sprintf("Mitigated %s", name))
}
