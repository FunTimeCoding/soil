package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, service_tester.WithInstances()).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
