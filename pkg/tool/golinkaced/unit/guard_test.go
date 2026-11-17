package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/links")
	v.VerifyGuarded(t, "/api/search")
	v.VerifyModelContext(t)
}
