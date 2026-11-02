package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestTheJobViewTakesTheRequestedJobAndItsTrace(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	c := clients[constant.FixtureInstance]
	c.SeedJob(1, "compile", "kilo", gitlabConstant.JobSuccess)
	c.SeedJob(2, "verify", "lima", gitlabConstant.JobFail)
	c.SeedTrace(1, "juliett")
	v, e := s.JobView(t.Context(), constant.FixtureInstance, 5, 10, 1)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, int(v.Selected.Identifier))
	assert.String(t, "juliett", v.Trace)
	assert.Count(t, 2, v.Jobs)
}

func TestTheJobViewPrefersARunningJobThenAFailedOne(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	c := clients[constant.FixtureInstance]
	c.SeedJob(1, "compile", "kilo", gitlabConstant.JobSuccess)
	c.SeedJob(2, "verify", "lima", gitlabConstant.JobFail)
	v, e := s.JobView(t.Context(), constant.FixtureInstance, 5, 10, 0)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, int(v.Selected.Identifier))
	c.SeedJob(3, "ship", "mike", gitlabConstant.JobRunning)
	w, f := s.JobView(t.Context(), constant.FixtureInstance, 5, 10, 0)
	assert.FatalOnError(t, f)
	assert.Integer(t, 3, int(w.Selected.Identifier))
}
