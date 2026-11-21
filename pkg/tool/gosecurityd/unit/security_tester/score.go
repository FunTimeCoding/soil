package security_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
)

func (o *Tester) Score() *score.Service {
	o.t.Helper()

	if o.score != nil {
		return o.score
	}

	g := peer.New(
		constant.NationalDatabasePeer,
		"national-database.test",
		o.NationalDatabase,
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.score = score.New(o.Store(), o.NationalDatabase, g, o.now)

	return o.score
}
