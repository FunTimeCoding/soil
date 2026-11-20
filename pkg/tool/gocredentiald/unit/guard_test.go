package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/database_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, database_tester.New(t).Write()).Served()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
