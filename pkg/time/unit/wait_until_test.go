package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"sync/atomic"
	"testing"
	"time"
)

func TestUntilWaitsForCondition(t *testing.T) {
	var polled atomic.Int32
	wait_tester.Until(t, func() bool { return polled.Add(1) >= 3 }, time.Second)
	assert.Integer(t, 3, int(polled.Load()))
}
