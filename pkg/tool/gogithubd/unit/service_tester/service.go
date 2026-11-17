package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		g := peer.New(
			constant.GitHubPeer,
			constant.GitHubAddress,
			o.Source,
			logger.New(o.t.Context()),
			o.Reporter,
		)
		o.t.Cleanup(g.Close)
		o.service = service.New(o.Source, g, "alfa")
	}

	return o.service
}
