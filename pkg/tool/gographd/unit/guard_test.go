package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
