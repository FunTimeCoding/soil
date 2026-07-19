package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/result/metric"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestMetric(t *testing.T) {
	assert.NotNil(t, metric.New(upper.Alfa))
}
