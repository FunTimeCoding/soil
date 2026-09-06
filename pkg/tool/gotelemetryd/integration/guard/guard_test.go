package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/events")
	c.VerifyGuarded(t, "/api/summary")
	c.VerifyOpenPost(t, "/api/events")
	c.VerifyOpen(t, constant.HeatmapPath)
	c.VerifyModelContext(t)
}
