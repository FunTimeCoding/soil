package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"testing"
)

func TestLokiBasicConstant(t *testing.T) {
	assert.String(t, "stderr", constant.Stderr)
}
