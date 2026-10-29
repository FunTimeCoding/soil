package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/client"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestTheUsageClientAndItsMockShareTheFace(t *testing.T) {
	c := client.New("nextcloud.test", 443, true, "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.UsageClient{c, m}
	assert.Count(t, 2, faces)
}

func TestARefusedUsageIsNotReachable(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusUnauthorized, "")
	host, port := upstream_tester.HostPort(t, s.URL)
	c := client.New(host, port, false, "token")
	defer c.Close()
	assert.Error(t, c.Reachable(t.Context()))
}
