package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (w *Web) Guard(
	m *http.ServeMux,
	r face.Reporter,
) (*guard.Mux, func()) {
	result := guard.New(m, w.ServiceTokens)

	if w.Authorization == nil {
		return result, func() {}
	}

	c := w.Authorization().WithReporter(r)
	c.Mount(result)

	return result, c.Close
}
