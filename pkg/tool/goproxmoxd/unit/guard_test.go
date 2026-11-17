package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	c := base.New(t).Served()
	c.VerifyBase(t)
	c.VerifyFavicon(t)
	c.VerifyOpen(t, constant.FloorPath)
	c.VerifyOpen(t, "/event")
	c.VerifyModelContext(t)
}
