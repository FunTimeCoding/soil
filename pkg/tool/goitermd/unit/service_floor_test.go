package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
	"testing"
	"time"
)

func TestSendKeysPacesAtLeastTheFloor(t *testing.T) {
	s := service.New(sessionClient("alfa"), 50*time.Millisecond)
	start := time.Now()
	assert.FatalOnError(t, s.SendKeys("alfa", []string{"up", "return"}, 0))
	assert.True(t, time.Since(start) >= 50*time.Millisecond)
}
