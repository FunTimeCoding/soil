package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic"
	"testing"
)

func TestMapForLabel(t *testing.T) {
	assert.Any(
		t,
		map[string][]*generic.Result{
			"differing-container1": {
				{
					Type:   constant.Scalar,
					Metric: `node_load1{container="differing-container1"}`,
					Value:  "1",
				},
			},
		},
		generic.MapForLabel(
			[]*generic.Result{
				{
					Type:   constant.Scalar,
					Metric: `node_load1{container="differing-container1"}`,
					Value:  "1",
				},
			},
			constant.ContainerLabel,
		),
	)
}
