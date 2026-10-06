package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/url"
	"strconv"
	"testing"
)

func HostPort(
	t *testing.T,
	server string,
) (string, int) {
	t.Helper()
	u, e := url.Parse(server)
	assert.FatalOnError(t, e)
	port, f := strconv.Atoi(u.Port())
	assert.FatalOnError(t, f)

	return u.Hostname(), port
}
