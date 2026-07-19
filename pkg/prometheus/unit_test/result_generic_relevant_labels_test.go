package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic"
	"testing"
)

func TestRelevantLabels(t *testing.T) {
	// TODO: Enter valid parseable results
	assert.Any(
		t,
		[]string{"container"},
		generic.RelevantLabels(
			[]*generic.Result{
				{
					Type:   constant.Scalar,
					Metric: `node_load1{container="differing-container1", pod="differing-pod1", namespace="identical-namespace"}`,
					Value:  "1",
				},
				{
					Type:   constant.Scalar,
					Metric: `node_load1{container="differing-container2", pod="differing-pod2", namespace="identical-namespace"}`,
					Value:  "2",
				},
			},
			[]string{"pod"},
		),
	)
}
