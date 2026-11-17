package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/trigger")
	v.VerifyModelContext(t)
}
