package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.Reports(t, t.TempDir())
	v.VerifyBase(t)
	v.VerifyFavicon(t)
	v.VerifySignIn(t)
}
