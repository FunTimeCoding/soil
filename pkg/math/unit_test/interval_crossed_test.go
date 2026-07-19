package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/interval_crossed"
	"testing"
)

func TestIntervalCrossedFloat(t *testing.T) {
	assert.False(t, interval_crossed.Float(10, 0, 1))
	assert.False(t, interval_crossed.Float(10, 1, 9))
	assert.True(t, interval_crossed.Float(10, 9, 10))
	assert.True(t, interval_crossed.Float(10, 9, 11))
	assert.False(t, interval_crossed.Float(10, 10, 11))
}
