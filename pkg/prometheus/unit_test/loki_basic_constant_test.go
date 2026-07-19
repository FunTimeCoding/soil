package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"testing"
)

func TestLokiBasicConstant(t *testing.T) {
	assert.String(t, constant.Stderr, "stderr")
}
