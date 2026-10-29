package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func refusedClient(
	t *testing.T,
	status int,
	answer string,
) *client.Client {
	t.Helper()
	s, _ := upstream_tester.New(t, status, answer)
	host, port := upstream_tester.HostPort(t, s.URL)
	result := client.New(locator.New(host).Insecure().Port(port), "token")
	t.Cleanup(result.Close)

	return result
}
