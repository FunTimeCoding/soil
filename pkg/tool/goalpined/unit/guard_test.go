package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := mountedUpload(t, service_tester.New(t).Service)
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/packages")
	v.VerifyModelContext(t)
}
