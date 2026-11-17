package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyOpen(t, "/jwks")
	v.VerifyOpenPost(t, "/token")
	v.VerifyFavicon(t)
	v.VerifyModelContext(t)
	v.VerifySignIn(t)
	v.VerifySession(t, webConstant.PalettePath)
}
