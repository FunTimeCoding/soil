package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	c := base.New(t).Served()
	c.VerifyBase(t)
	c.VerifyModelContext(t)
}
