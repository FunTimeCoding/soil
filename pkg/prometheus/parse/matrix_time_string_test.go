package parse

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/prometheus/common/model"
	"testing"
)

func TestMatrixTimeString(t *testing.T) {
	assert.Count(
		t,
		1,
		MatrixTimeString(
			model.Matrix{
				{
					Metric:     model.Metric{},
					Values:     []model.SamplePair{{Timestamp: 0, Value: 1}},
					Histograms: []model.SampleHistogramPair{},
				},
			},
		),
	)
}
