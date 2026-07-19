package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/metric"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "METRIC_PORT", metric.PortEnvironment)
}
