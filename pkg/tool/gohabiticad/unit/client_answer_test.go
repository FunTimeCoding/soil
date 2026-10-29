package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestTheClientReadsTasksAndTags(t *testing.T) {
	m := service_tester.NewSource(t)
	daily := task.Stub()
	daily.Identifier = "task-1"
	daily.Text = "Deploy fleet"
	daily.Type = "dailys"
	m.AddTask(daily)
	habit := task.Stub()
	habit.Identifier = "task-2"
	habit.Text = "Stretch"
	habit.Type = "habit"
	m.AddTask(habit)
	g := tag.Stub()
	g.Identifier = "tag-1"
	g.Name = "deploy"
	m.AddTag(g)
	c := client_tester.New(
		t,
		service_tester.Service(t, m),
		constant.ModelContextTestToken,
	)
	all, e := c.GetTasks(t.Context(), "")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, all)
	dailies, h := c.GetTasks(t.Context(), "dailys")
	assert.FatalOnError(t, h)
	assert.Count(t, 1, dailies)
	assert.String(t, "Deploy fleet", dailies[0].Text)
	tags, f := c.GetTags(t.Context())
	assert.FatalOnError(t, f)
	assert.Count(t, 1, tags)
	assert.String(t, "deploy", tags[0].Name)
}

func TestTheClientCreatesATask(t *testing.T) {
	c := client_tester.New(
		t,
		service_tester.Service(t, service_tester.NewSource(t)),
		constant.ModelContextTestToken,
	)
	v, e := c.CreateTask(t.Context(), "todo", "Ship loop C", "")
	assert.FatalOnError(t, e)
	assert.String(t, "Ship loop C", v.Text)
}

func TestTheClientAnswersARefusedStatAsValidation(t *testing.T) {
	m := service_tester.NewSource(t)
	m.Failure = validation.New("%q is not a valid Stat.", "per")
	c := client_tester.New(
		t,
		service_tester.Service(t, m),
		constant.ModelContextTestToken,
	)
	_, e := c.AllocateStat(t.Context(), "per")
	assert.True(t, validation.Is(e))
}

func TestTheClientReachesTheDaemon(t *testing.T) {
	c := client_tester.New(
		t,
		service_tester.Service(t, service_tester.NewSource(t)),
		constant.ModelContextTestToken,
	)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestTheClientRefusesAWrongToken(t *testing.T) {
	c := client_tester.New(
		t,
		service_tester.Service(t, service_tester.NewSource(t)),
		"wrong",
	)
	assert.Error(t, c.Reachable(t.Context()))
	_, e := c.GetTags(t.Context())
	assert.Error(t, e)
}
