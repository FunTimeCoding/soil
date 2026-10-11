package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	c := s.Server
	c.VerifyBase(t)
	c.VerifyInterface(t)
	c.VerifyOpen(t, constant.FloorPath)
	c.VerifyOpen(t, "/event")
	c.VerifyModelContext(t)
}
