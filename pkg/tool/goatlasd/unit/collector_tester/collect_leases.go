package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client/mock_client"
	"testing"
)

func CollectLeases(
	t *testing.T,
	f *mock_client.Client,
) []*sighting.Sighting {
	t.Helper()
	result, e := lease.New(f, peer_tester.NewGuard(t, f)).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)

	return result
}
