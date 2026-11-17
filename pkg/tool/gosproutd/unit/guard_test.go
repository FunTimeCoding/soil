package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyFavicon(t)
	v.VerifyOpen(t, constant.DashboardPath)
	v.VerifyOpen(t, constant.SessionsPath)
	v.VerifyOpen(t, constant.CountersPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
