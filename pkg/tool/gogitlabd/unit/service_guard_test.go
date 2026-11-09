package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestUnreachableInstanceFailsFastWithoutBeingAskedAgain(t *testing.T) {
	s, clients := service_tester.New(
		t,
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)
	alfa := clients[constant.FixtureInstance]
	alfa.FailProjects(refusedConnection())
	_, e := s.Projects(t.Context(), constant.FixtureInstance, "")
	assert.True(t, unavailable.Is(e))
	alfa.FailProjects(nil)
	_, f := s.Projects(t.Context(), constant.FixtureInstance, "")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "gitlab:alfa unreachable since", f.Error())
	_, g := s.Projects(t.Context(), constant.FixtureOtherInstance, "")
	assert.FatalOnError(t, g)
}
