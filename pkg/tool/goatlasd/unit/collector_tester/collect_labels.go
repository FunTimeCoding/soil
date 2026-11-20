package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
)

func (o *Tester) CollectLabels(f *mock_client.Client) []*placement.Placement {
	o.t.Helper()
	result, e := label.New(f, peer_tester.NewGuard(o.t, f)).Collect(
		context.Background(),
		nil,
	)
	assert.FatalOnError(o.t, e)

	return result
}
