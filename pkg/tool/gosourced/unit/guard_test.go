package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t, service_tester.ServiceTestdata("cross-package/src"))
	defer s.Stop()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyModelContext(t)
}
