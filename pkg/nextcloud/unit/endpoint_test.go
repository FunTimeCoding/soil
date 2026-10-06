package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/client"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestPropfindAcceptsMultiStatus(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusMultiStatus,
		"<d:multistatus/>",
	)
	assert.FatalOnError(t, newBasic(t, s.URL).Propfind())
	r := last()
	assert.String(t, "PROPFIND", r.Method)
	assert.String(t, "/remote.php/dav/files/alfa/", r.URL.Path)
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo-password", password)
}

func TestPropfindRefusalIsUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusUnauthorized, "")
	e := newBasic(t, s.URL).Propfind()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 401", e.Error())
}

func TestUsageSendsTheTokenAndReadsCounts(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"ocs":{"data":{"nextcloud":{"storage":{"num_files":12},"shares":{"num_shares":3}}}}}`,
	)
	host, p := upstream_tester.HostPort(t, s.URL)
	u, e := client.New(host, p, false, "charlie-token").Fetch()
	assert.FatalOnError(t, e)
	assert.Any(t, int64(12), u.Files)
	assert.Any(t, int64(3), u.Shares)
	r := last()
	assert.String(t, "/ocs/v2.php/apps/serverinfo/api/v1/info", r.URL.Path)
	assert.String(t, "json", r.URL.Query().Get(constant.FormatParameter))
	assert.String(t, "charlie-token", r.Header.Get(constant.TokenHeader))
}

func TestUsageRefusalIsUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusForbidden, "")
	host, p := upstream_tester.HostPort(t, s.URL)
	_, e := client.New(host, p, false, "charlie-token").Fetch()
	assert.True(t, unexpected.Is(e))
}
