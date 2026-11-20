package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
	v.VerifyFavicon(t)
	v.VerifyOpen(t, constant.FlagsPath)
	v.VerifyOpen(t, constant.LookupPath)
	v.VerifyOpen(t, constant.SystemsPath)
	v.VerifyOpen(t, constant.SystemPath)
	v.VerifyOpenPost(t, constant.FlagPath)
	v.VerifyOpenPost(t, constant.MitigatePath)
}
