package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyInterface(t)
	c.VerifyModelContext(t)
}
