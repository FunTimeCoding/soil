package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	c := s.Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, "/api/certificates")
	c.VerifyOpen(t, "/favicon.ico")
	c.VerifyModelContext(t)
}
