package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.ContextServer
	c.VerifyBase(t)
	c.VerifyInterface(t)
	c.VerifyGuarded(t, "/api/libraries")
	c.VerifyGuarded(t, "/api/sessions")
	c.VerifyModelContext(t)
}
