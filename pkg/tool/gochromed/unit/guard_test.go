package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
