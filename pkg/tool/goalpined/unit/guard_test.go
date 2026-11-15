package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, service_tester.New(t).Option)
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/packages")
	v.VerifyModelContext(t)
}
