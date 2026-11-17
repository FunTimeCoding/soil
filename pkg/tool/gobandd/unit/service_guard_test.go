package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestUnreachableHostFailsFastWithoutBeingAskedAgain(t *testing.T) {
	o := service_tester.New(t, service_tester.WithInstances("alfa", "bravo"))
	v := o.Service()
	alfa := o.Alfa()
	alfa.Fail(outage_tester.RefusedDial())
	_, e := v.Status("alfa")
	assert.True(t, unavailable.Is(e))
	alfa.Fail(nil)
	_, f := v.Status("alfa")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "band:alfa unreachable since", f.Error())
	_, g := v.Status("bravo")
	assert.FatalOnError(t, g)
}
