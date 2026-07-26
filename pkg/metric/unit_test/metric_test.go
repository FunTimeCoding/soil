package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/metric/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "METRIC_PORT", constant.PortEnvironment)
}
