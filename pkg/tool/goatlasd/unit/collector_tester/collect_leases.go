package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_lease_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"testing"
)

func CollectLeases(
	t *testing.T,
	f *mock_lease_source.Client,
) []*sighting.Sighting {
	t.Helper()
	result, e := lease.New(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)

	return result
}
