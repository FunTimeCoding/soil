package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_inventory_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"testing"
)

func CollectLabels(
	t *testing.T,
	f *mock_inventory_source.Client,
) []*placement.Placement {
	t.Helper()
	result, e := label.New(f).Collect(context.Background(), nil)
	assert.FatalOnError(t, e)

	return result
}
