package security_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/feed"
)

func (o *Tester) Feed(a *mock_client.Client) *feed.Service {
	o.t.Helper()
	l := logger.New(o.t.Context())
	g := peer.New(constant.AnsiblePeer, a.Address(), a, l, o.Reporter)
	o.t.Cleanup(g.Close)

	return feed.New(a, g, o.System(), l)
}
