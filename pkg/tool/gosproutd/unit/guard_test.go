package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"testing"
)

func TestGuard(t *testing.T) {
	v, _ := webServer(t)
	v.VerifyBase(t)
	v.VerifyFavicon(t)
	v.VerifyOpen(t, constant.DashboardPath)
	v.VerifyOpen(t, constant.SessionsPath)
	v.VerifyOpen(t, constant.CountersPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
