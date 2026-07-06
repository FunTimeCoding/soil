package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func ServeAsynchronous(s *http.Server) {
	go func() {
		errors.PanicOnUnclean(s.ListenAndServe())
	}()
}
