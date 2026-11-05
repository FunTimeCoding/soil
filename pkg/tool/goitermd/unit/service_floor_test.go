package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/service_tester"
	"testing"
	"time"
)

func TestSendKeysPacesAtLeastTheFloor(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	s := service_tester.NewPaced(t, c, 50*time.Millisecond)
	start := time.Now()
	assert.FatalOnError(t, s.SendKeys("alfa", []string{"up", "return"}, 0))
	assert.True(t, time.Since(start) >= 50*time.Millisecond)
}
