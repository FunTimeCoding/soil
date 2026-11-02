package unit

import (
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func closedLocator(t *testing.T) *locator.Locator {
	t.Helper()
	s := httptest.NewServer(http.NotFoundHandler())
	host, port := network.SplitHostPort(s.Listener.Addr().String())
	s.Close()

	return locator.New(host).Port(port).Insecure()
}
