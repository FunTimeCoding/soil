package collector_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/kubernetes"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_pod_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"testing"
)

func CollectPods(
	t *testing.T,
	f *mock_pod_source.Client,
) []*placement.Placement {
	t.Helper()
	result, e := kubernetes.New(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)

	return result
}
