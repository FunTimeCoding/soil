package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyInterface(t)
	c.VerifyOpen(t, constant.FloorPath)
	c.VerifyOpen(t, "/event")
	c.VerifyModelContext(t)
}
