package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/integration/base"
	"testing"
)

func TestGuard(t *testing.T) {
	c := base.New(t).Server
	c.VerifyBase(t)
	c.VerifyGuarded(t, constant.ClaimsPath)
	c.VerifyGuarded(t, constant.StreamPath)
}
