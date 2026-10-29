package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
	"testing"
	"time"
)

func TestSendKeysPacesAtLeastTheFloor(t *testing.T) {
	c := sessionClient("alfa")
	o := option.New()
	o.Source = func() face.ItermSource { return c }
	o.KeyInterval = 50 * time.Millisecond
	s := service.New(o)
	t.Cleanup(s.Close)
	start := time.Now()
	assert.FatalOnError(t, s.SendKeys("alfa", []string{"up", "return"}, 0))
	assert.True(t, time.Since(start) >= 50*time.Millisecond)
}
