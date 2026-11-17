package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/client"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestUsageClientAndItsMockShareFace(t *testing.T) {
	c := client.New("nextcloud.test", 443, true, "token")
	defer c.Close()
	m := mock_client.New(t)
	faces := []face.UsageClient{c, m}
	assert.Count(t, 2, faces)
}

func TestRefusedUsageIsNotReachable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
	)
	host, port := upstream_server.HostPort(t, s.URL)
	c := client.New(host, port, false, "token")
	defer c.Close()
	assert.Error(t, c.Reachable(t.Context()))
}
