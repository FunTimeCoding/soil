package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/constant"
	"testing"
)

func TestLokiConstant(t *testing.T) {
	assert.String(t, "LOKI_HOST", constant.HostEnvironment)
}
