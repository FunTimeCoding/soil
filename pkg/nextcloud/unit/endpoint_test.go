package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/nextcloud/unit/nextcloud_tester"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/client"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestPropfindAcceptsMultiStatus(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusMultiStatus),
		upstream_server.WithAnswer("<d:multistatus/>"),
	)
	assert.FatalOnError(t, nextcloud_tester.New(t, s.URL).Propfind())
	r := s.Last()
	assert.String(t, "PROPFIND", r.Method)
	assert.String(t, "/remote.php/dav/files/alfa/", r.URL.Path)
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo-password", password)
}

func TestPropfindRefusalIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
	)
	e := nextcloud_tester.New(t, s.URL).Propfind()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 401", e.Error())
}

func TestUsageSendsTokenAndReadsCounts(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"ocs":{"data":{"nextcloud":{"storage":{"num_files":12},"shares":{"num_shares":3}}}}}`,
		))
	host, p := upstream_server.HostPort(t, s.URL)
	owned := client.New(host, p, false, "charlie-token")
	defer owned.Close()
	u, e := owned.Fetch()
	assert.FatalOnError(t, e)
	assert.Any(t, int64(12), u.Files)
	assert.Any(t, int64(3), u.Shares)
	r := s.Last()
	assert.String(t, "/ocs/v2.php/apps/serverinfo/api/v1/info", r.URL.Path)
	assert.String(t, "json", r.URL.Query().Get(constant.FormatParameter))
	assert.String(t, "charlie-token", r.Header.Get(constant.TokenHeader))
}

func TestUsageRefusalIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
	)
	host, p := upstream_server.HostPort(t, s.URL)
	owned := client.New(host, p, false, "charlie-token")
	defer owned.Close()
	_, e := owned.Fetch()
	assert.True(t, unexpected.Is(e))
}
