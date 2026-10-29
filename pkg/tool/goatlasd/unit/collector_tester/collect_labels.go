package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func CollectLabels(
	t *testing.T,
	f *mock_client.Client,
) []*placement.Placement {
	t.Helper()
	result, e := label.New(f, peer_tester.NewGuard(t, f)).Collect(
		context.Background(),
		nil,
	)
	assert.FatalOnError(t, e)

	return result
}
