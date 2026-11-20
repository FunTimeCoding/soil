package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client/mock_client"
)

func (o *Tester) CollectLeases(f *mock_client.Client) []*sighting.Sighting {
	o.t.Helper()
	result, e := lease.New(f, peer_tester.NewGuard(o.t, f)).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(o.t, e)

	return result
}
