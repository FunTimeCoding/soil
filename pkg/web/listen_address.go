package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func ListenAddress(
	m *http.ServeMux,
	address string,
) {
	s := Server(m, address)
	defer errors.PanicClose(s)
	errors.PanicOnError(s.ListenAndServe())
}
