package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"testing"
)

func TestUnreachableHostFailsFastWithoutBeingAskedAgain(t *testing.T) {
	v, engines := service_tester.New(t, memory.New(), "alfa", "bravo")
	alfa := engines["alfa"]
	alfa.Fail(refusedConnection())
	_, e := v.Status("alfa")
	assert.True(t, unavailable.Is(e))
	alfa.Fail(nil)
	_, f := v.Status("alfa")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "band:alfa unreachable since", f.Error())
	_, g := v.Status("bravo")
	assert.FatalOnError(t, g)
}
