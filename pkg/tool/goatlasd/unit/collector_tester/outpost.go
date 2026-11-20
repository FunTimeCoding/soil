package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/outpost"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
)

func (o *Tester) Outpost(clients ...*mock_client.Client) *outpost.Collector {
	o.t.Helper()
	var targets []*target.Target

	for i, c := range clients {
		targets = append(
			targets,
			target.New(
				constant.FixtureOutpostNames[i],
				c,
				peer_tester.NewGuard(o.t, c),
			),
		)
	}

	return outpost.New(targets, logger.New(context.Background()))
}
