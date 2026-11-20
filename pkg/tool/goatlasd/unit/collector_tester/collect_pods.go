package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/kubernetes"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
)

func (o *Tester) CollectPods(f *mock_client.Client) []*placement.Placement {
	o.t.Helper()
	result, e := kubernetes.New(f, peer_tester.NewGuard(o.t, f)).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(o.t, e)

	return result
}
