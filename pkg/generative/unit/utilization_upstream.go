package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func utilizationUpstream(
	t *testing.T,
	answer http.HandlerFunc,
) *utilization.Client {
	t.Helper()
	s := httptest.NewServer(answer)
	t.Cleanup(s.Close)
	host, port := network.SplitHostPort(s.Listener.Addr().String())
	result := utilization.New(locator.New(host).Port(port).Insecure())
	t.Cleanup(result.Close)

	return result
}
