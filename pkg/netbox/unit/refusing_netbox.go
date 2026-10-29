package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func refusingNetbox(
	t *testing.T,
	status int,
	body string,
) *netbox.Client {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				_, e := w.Write([]byte(body))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	result := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	t.Cleanup(result.Close)

	return result
}
