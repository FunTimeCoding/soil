package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"strconv"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	d := lite.NewMemory()
	o.t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(
		constant.ForgePeer,
		"",
		o.Forge,
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(
		store.New(d),
		g,
		publish.New(
			o.Forge,
			strconv.Itoa(constant.FixtureProject),
			constant.PublishBranch,
			constant.FixtureAuthorityDirectory,
			constant.FixtureClusterAuthority,
			constant.FixtureSecretPath,
		),
		o.Clock,
	)

	return o.service
}
