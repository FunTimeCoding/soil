package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	c := s.Server
	c.VerifyBase(t)
	c.VerifyModelContext(t)
}
