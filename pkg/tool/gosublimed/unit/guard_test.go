package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	c := s.Server
	c.VerifyBase(t)
	c.VerifyInterface(t)
	c.VerifyModelContext(t)
}
