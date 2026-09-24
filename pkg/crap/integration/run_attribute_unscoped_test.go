package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"testing"
)

func TestRunAttributeRefusesUnscoped(t *testing.T) {
	o := option.NewAttribute()
	o.Root = module(t)
	o.Patterns = []string{constant.AllPackages}
	out := assert.Capture(
		t,
		func() { assert.Integer(t, 1, crap.RunAttribute(o)) },
	)
	assert.StringContains(t, "name the packages", out)
}

func TestRunAttributeScopedPasses(t *testing.T) {
	o := option.NewAttribute()
	o.Root = module(t)
	o.Patterns = []string{"./pkg/a/..."}
	o.Notation = true
	assert.Capture(t, func() { assert.Integer(t, 0, crap.RunAttribute(o)) })
}
