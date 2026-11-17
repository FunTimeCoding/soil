package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, service_tester.WithInstances()).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
