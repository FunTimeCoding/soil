package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (s *Server) createRaid(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	filenames := r.Form["fileNames"]
	identifier := s.store.CreateRaid(filenames)
	http.Redirect(
		w,
		r,
		fmt.Sprintf("/raids/%d", identifier),
		http.StatusSeeOther,
	)
}
