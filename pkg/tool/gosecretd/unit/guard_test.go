package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.ContextServer
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/status")
	c.VerifyGuarded(t, "/api/secrets")
	c.VerifyModelContext(t)
}
