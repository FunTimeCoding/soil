package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/jellyfin"
	"net"
	"net/url"
	"strconv"
	"testing"
)

func newJellyfin(
	t *testing.T,
	server string,
) *jellyfin.Client {
	t.Helper()
	u, e := url.Parse(server)
	assert.FatalOnError(t, e)
	host, port, f := net.SplitHostPort(u.Host)
	assert.FatalOnError(t, f)
	number, g := strconv.Atoi(port)
	assert.FatalOnError(t, g)
	result := jellyfin.New(host, number, "alfa-token")
	t.Cleanup(result.Close)

	return result
}
