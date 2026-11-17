package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestQuietRefusalIsNotCapturedByGuard(t *testing.T) {
	s := service_tester.New(t)
	s.Source.Failure = validation.New("%q is not a valid Stat.", "bravo")
	v := s.Service()
	_, e := v.Allocate("bravo")
	assert.True(t, validation.Is(e))
	s.Source.Failure = not_found.New("gear", "weapon_missing")
	_, f := v.Equip("weapon_missing")
	assert.True(t, not_found.Is(f))
	assert.Count(t, 0, s.Reporter.Events())
}
