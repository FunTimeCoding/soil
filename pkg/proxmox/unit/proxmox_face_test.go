package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/proxmox"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"testing"
)

func TestTheClientAndItsMockKeepTheProxmoxFace(t *testing.T) {
	c := proxmox.New("alfa.example", proxmox.WithToken("alfa", "secret"))
	t.Cleanup(c.Close)
	m := mock_client.New()
	clients := []face.ProxmoxClient{c, m}
	m.Close()
	assert.Count(t, 2, clients)
	assert.Integer(t, 1, m.Closed)
}
