package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	c := base.New(t).Served()
	c.VerifyBase(t)
	c.VerifyFavicon(t)
	c.VerifyGuarded(t, "/api/events")
	c.VerifyGuarded(t, "/api/summary")
	c.VerifyOpenPost(t, "/api/events")
	c.VerifyOpen(t, constant.HeatmapPath)
	c.VerifyModelContext(t)
}
