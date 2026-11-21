package security_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
)

func (o *Tester) FreshScore() *score.Service {
	o.t.Helper()
	g := peer.New(
		constant.NationalDatabasePeer,
		"national-database.test",
		o.NationalDatabase,
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)

	return score.New(o.Store(), o.NationalDatabase, g, o.now)
}
