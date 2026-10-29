package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/outpost"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"testing"
)

func outpostCollector(
	t *testing.T,
	clients ...*mock_client.Client,
) *outpost.Collector {
	var targets []*target.Target

	for i, c := range clients {
		targets = append(
			targets,
			target.New(
				constant.FixtureOutpostNames[i],
				c,
				peer_tester.NewGuard(t, c),
			),
		)
	}

	return outpost.New(targets, logger.New(context.Background()))
}
