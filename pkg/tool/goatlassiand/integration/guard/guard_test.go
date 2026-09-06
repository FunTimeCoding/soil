package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyInterface(t)
	c.VerifyOpen(t, constant.PlatePath)
	c.VerifyOpen(t, "/event")
	c.VerifyModelContext(t)
}
