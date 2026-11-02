package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func TestGuard(t *testing.T) {
	v := mounted(t, fixtureDatabase(t), memory.New())
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
