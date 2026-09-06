package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t, "../../service/testdata/cross-package/src")
	defer s.Stop()
	c := s.Server
	c.VerifyBase(t)
	c.VerifyModelContext(t)
}
