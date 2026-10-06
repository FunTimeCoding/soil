package unit

import (
	"github.com/funtimecoding/soil/pkg/argocd"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newClient(
	t *testing.T,
	server string,
) *argocd.Client {
	t.Helper()
	host, port := upstream_tester.HostPort(t, server)

	return argocd.New(host, port, false, "alfa-token")
}
