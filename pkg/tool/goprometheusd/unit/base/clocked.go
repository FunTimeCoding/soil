package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/service_tester"
	"testing"
)

func Clocked(t *testing.T) (*model_context_client.Client, *mock_client.Client) {
	t.Helper()
	s := New(
		t,
		service_tester.WithInstances("alfa", "bravo"),
		service_tester.WithClock(clock_tester.Window),
	)
	result := s.Tools()
	result.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "alfa"},
	)

	return result, s.Alfa()
}
