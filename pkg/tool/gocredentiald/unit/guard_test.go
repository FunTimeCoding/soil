package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/database_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t, database_tester.Write(t), memory.New())
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
