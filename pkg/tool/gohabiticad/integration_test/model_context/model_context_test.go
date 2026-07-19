package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/habitica/tag"
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	c := o.Client
	assert.Count(t, 9, c.ListTools())
	k := task.Stub()
	k.ID = "task-1"
	k.Text = "Deploy fleet"
	k.Type = "daily"
	o.MockClient.AddTask(k)
	g := tag.Stub()
	g.ID = "tag-1"
	g.Name = "deploy"
	o.MockClient.AddTag(g)
	tasks := c.MustCallTool(constant.GetTasks, map[string]any{})
	assert.StringContains(t, "Deploy fleet", tasks)
	tags := c.MustCallTool(constant.GetTags, map[string]any{})
	assert.StringContains(t, "deploy", tags)
	stats := c.MustCallTool(constant.GetStats, map[string]any{})
	assert.StringContains(t, "warrior", stats)
	score := c.MustCallTool(
		constant.ScoreTask,
		map[string]any{
			parameter.Identifier: "task-1",
			constant.Direction:   "up",
		},
	)
	assert.StringContains(t, "hp", score)
	gear := c.MustCallTool(constant.GetGear, map[string]any{})
	assert.StringContains(t, "equipped", gear)
	cron := c.MustCallTool(constant.Cron, map[string]any{})
	assert.StringContains(t, "rolled_over", cron)
}
