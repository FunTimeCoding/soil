package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestTokenSummaryCountsWithheldMemories(t *testing.T) {
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
	s, g := c.GetTokens(x, &client.GetTokensParams{})
	assert.FatalOnError(t, g)
	assert.Count(t, 1, s.Statistic)
	assert.Integer(t, 1, s.Withheld)
	assert.Integer(t, 2, s.Memories)
}
