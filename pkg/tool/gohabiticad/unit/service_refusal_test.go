package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestAQuietRefusalIsNotCapturedByTheGuard(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = validation.New("%q is not a valid Stat.", "bravo")
	m := memory.New()
	v := service_tester.Shared(t, c, m)
	_, e := v.Allocate("bravo")
	assert.True(t, validation.Is(e))
	c.Failure = not_found.New("gear", "weapon_missing")
	_, f := v.Equip("weapon_missing")
	assert.True(t, not_found.Is(f))
	assert.Count(t, 0, m.Events())
}
