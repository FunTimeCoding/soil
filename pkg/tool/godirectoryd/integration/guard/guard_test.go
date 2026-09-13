package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/user")
	c.VerifyGuarded(t, "/api/group")
	c.VerifyModelContext(t)
}
