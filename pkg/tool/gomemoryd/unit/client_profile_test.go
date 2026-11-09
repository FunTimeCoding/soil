package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestProfileLeavesHiddenMemoriesOut(t *testing.T) {
	c := base.NewHidden(t, "private").MemoryClient(t)
	x := t.Context()
	_, e := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:        "alfa",
			Content:     "alfa content",
			Description: "alfa",
		},
	)
	assert.FatalOnError(t, e)
	_, f := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:        "bravo",
			Content:     "bravo content",
			Description: "bravo",
			Tags:        &[]string{"private"},
		},
	)
	assert.FatalOnError(t, f)
	p, g := c.GetProfile(x, &client.GetProfileParams{})
	assert.FatalOnError(t, g)
	assert.Count(t, 1, p.Index)
	assert.String(t, "alfa", p.Index[0].Name)
}
