package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
