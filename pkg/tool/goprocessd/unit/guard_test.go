package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, "alfa: sleep 60\n").Served()
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
	v.VerifyGuarded(t, "/api/processes")
	v.VerifyGuarded(t, "/api/processes/alfa/log")
}
