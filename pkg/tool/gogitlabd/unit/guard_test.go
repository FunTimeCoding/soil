package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	r := memory.New()
	d, _ := daemon_tester.New(t, r, constant.FixtureInstance)
	v := startServer(t, d, r)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyFavicon(t)
	v.VerifyOpen(t, constant.BoardPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
