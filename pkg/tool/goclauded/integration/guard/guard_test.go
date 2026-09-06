package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/sessions")
	c.VerifyOpen(t, constant.DashboardPath)
	c.VerifyOpen(t, "/event")
	c.VerifyModelContext(t)
}
