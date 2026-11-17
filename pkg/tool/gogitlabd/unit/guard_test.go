package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyFavicon(t)
	v.VerifyOpen(t, constant.BoardPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
