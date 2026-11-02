package unit

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newArgocdClient(
	t *testing.T,
	server string,
) *argocd.Client {
	t.Helper()
	host, port := upstream_tester.HostPort(t, server)
	result := argocd.New(host, port, false, "alfa-token")
	t.Cleanup(result.Close)

	return result
}
