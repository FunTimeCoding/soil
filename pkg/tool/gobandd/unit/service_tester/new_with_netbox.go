package service_tester

import (
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"testing"
)

func NewWithNetbox(
	t *testing.T,
	r face.Reporter,
	names ...string,
) (*service.Service, map[string]*mock_client.Client, *netbox.Client) {
	t.Helper()
	s, engines := New(t, r, names...)
	n := netbox.New()
	t.Cleanup(n.Close)
	g := peer.New(
		constant.Identity.Name(),
		n.Address(),
		n,
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(g.Close)

	return s.WithNetbox(n, g), engines, n
}
