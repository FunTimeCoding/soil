package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/service_tester"
	"testing"
	"time"
)

func TestSendKeysPacesAtLeastFloor(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithKeyInterval(50*time.Millisecond),
	)
	o.AddSession("alfa")
	start := time.Now()
	assert.FatalOnError(
		t,
		o.Service().SendKeys("alfa", []string{"up", "return"}, 0),
	)
	assert.True(t, time.Since(start) >= 50*time.Millisecond)
}
