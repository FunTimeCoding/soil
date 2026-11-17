package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/packages")
	v.VerifyModelContext(t)
}
