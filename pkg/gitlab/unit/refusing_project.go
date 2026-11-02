package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func refusingProject(
	t *testing.T,
	status int,
	message string,
) *client_tester.Tester {
	t.Helper()

	return client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					web.ObjectHeader(w)
					w.WriteHeader(status)
					web.Encode(w, map[string]string{"message": message})
				},
			)
		},
	)
}
