package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	c := netbox.New(locator.New(constant.Localhost), "token")
	t.Cleanup(c.Close)
	m := mock_client.New()
	t.Cleanup(m.Close)
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}
