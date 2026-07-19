package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/push"
	"testing"
)

func TestPushConstant(t *testing.T) {
	assert.Integer(t, 9091, push.Port)
}
