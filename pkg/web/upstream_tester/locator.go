package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"strings"
	"testing"
)

func Locator(
	t *testing.T,
	server string,
) *locator.Locator {
	t.Helper()
	host, port := HostPort(t, server)
	result := locator.New(host).Port(port)

	if strings.HasPrefix(server, constant.InsecurePrefix) {
		result.Insecure()
	}

	return result
}
