package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"testing"
)

func TestModelContext(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.Count(t, 9, c.ListTools())
	k := task.Stub()
	k.Identifier = "task-1"
	k.Text = "Deploy fleet"
	k.Type = "daily"
	s.Source.AddTask(k)
	g := tag.Stub()
	g.Identifier = "tag-1"
	g.Name = "deploy"
	s.Source.AddTag(g)
	tasks := c.MustCallTool(constant.GetTasks, map[string]any{})
	assert.StringContains(t, "Deploy fleet", tasks)
	tags := c.MustCallTool(constant.GetTags, map[string]any{})
	assert.StringContains(t, "deploy", tags)
	stats := c.MustCallTool(constant.GetStats, map[string]any{})
	assert.StringContains(t, "warrior", stats)
	score := c.MustCallTool(
		constant.ScoreTask,
		map[string]any{
			generativeConstant.ParameterIdentifier: "task-1",
			constant.Direction:                     "up",
		},
	)
	assert.StringContains(t, "hp", score)
	gear := c.MustCallTool(constant.GetGear, map[string]any{})
	assert.StringContains(t, "equipped", gear)
	cron := c.MustCallTool(constant.Cron, map[string]any{})
	assert.StringContains(t, "rolled_over", cron)
}
