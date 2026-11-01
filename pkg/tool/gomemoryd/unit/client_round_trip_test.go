package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"strings"
	"testing"
	"time"
)

func TestTheClientRoundTripsAMemory(t *testing.T) {
	c := base.New(t).MemoryClient(t)
	x := t.Context()
	scope := "alfa"
	file := "alfa.md"
	created, e := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:           "retry policy",
			Content:        "retry with backoff",
			Description:    "retry policy",
			Scope:          &scope,
			ProvenanceFile: &file,
		},
	)
	assert.FatalOnError(t, e)
	sourced, f := c.GetSourcedMemories(
		x,
		&client.GetSourcedMemoriesParams{Scope: scope},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, sourced)
	assert.Integer(t, int(created.Identifier), int(sourced[0].Identifier))
	content := "retry with jitter"
	_, g := c.PutMemory(
		x,
		created.Identifier,
		client.PutMemoryJSONRequestBody{Content: &content},
	)
	assert.FatalOnError(t, g)
	versions, h := c.GetVersions(x, "2000-01-01T00:00:00Z", 10)
	assert.FatalOnError(t, h)
	assert.Count(t, 2, versions)
	_, i := c.DeleteMemory(x, created.Identifier, &client.DeleteMemoryParams{})
	assert.FatalOnError(t, i)
	gone, j := c.GetSourcedMemories(
		x,
		&client.GetSourcedMemoriesParams{Scope: scope},
	)
	assert.FatalOnError(t, j)
	assert.Count(t, 0, gone)
}

func TestTheClientRoundTripsARelation(t *testing.T) {
	c := base.New(t).MemoryClient(t)
	x := t.Context()
	first, e := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:        "first",
			Content:     "alfa",
			Description: "first",
		},
	)
	assert.FatalOnError(t, e)
	second, f := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:        "second",
			Content:     "bravo",
			Description: "second",
		},
	)
	assert.FatalOnError(t, f)
	assert.FatalOnError(
		t,
		c.PostRelation(
			x,
			client.PostRelationJSONRequestBody{
				SourceIdentifier: first.Identifier,
				TargetIdentifier: second.Identifier,
			},
		),
	)
	relations, g := c.GetRelations(x, &client.GetRelationsParams{})
	assert.FatalOnError(t, g)
	assert.Count(t, 1, relations)
	assert.String(t, "first", relations[0].SourceName)
	assert.FatalOnError(
		t,
		c.DeleteRelation(
			x,
			&client.DeleteRelationParams{
				SourceIdentifier: first.Identifier,
				TargetIdentifier: second.Identifier,
			},
		),
	)
	left, h := c.GetRelations(x, &client.GetRelationsParams{})
	assert.FatalOnError(t, h)
	assert.Count(t, 0, left)
}

func TestTheClientReadsProfileTokensAndStatistics(t *testing.T) {
	c := base.New(t).MemoryClient(t)
	x := t.Context()
	_, e := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:        "retry policy",
			Content:     "retry with backoff",
			Description: "retry policy",
		},
	)
	assert.FatalOnError(t, e)
	p, f := c.GetProfile(x, &client.GetProfileParams{})
	assert.FatalOnError(t, f)
	assert.True(t, strings.Contains(p.Text, "retry policy"))
	tokens, g := c.GetTokens(x, &client.GetTokensParams{})
	assert.FatalOnError(t, g)
	assert.Count(t, 1, tokens.Statistic)
	_, h := c.GetStatistics(x)
	assert.FatalOnError(t, h)
	redacted, i := c.GetRedactedMemories(x)
	assert.FatalOnError(t, i)
	assert.Count(t, 0, redacted)
	assert.FatalOnError(t, c.PostImpressions(x, "held", "test", time.Now()))
}

func TestTheClientClassifiesRefusals(t *testing.T) {
	c := base.New(t).MemoryClient(t)
	x := t.Context()
	_, e := c.DeleteMemory(x, 9999, &client.DeleteMemoryParams{})
	assert.True(t, not_found.Is(e))
	content := "x"
	_, f := c.PutMemory(
		x,
		9999,
		client.PutMemoryJSONRequestBody{Content: &content},
	)
	assert.True(t, not_found.Is(f))
	_, g := c.PostMemory(
		x,
		client.PostMemoryJSONRequestBody{
			Name:        "x",
			Content:     "x",
			Description: "two\nlines",
		},
	)
	assert.True(t, validation.Is(g))
	h := c.DeleteRelation(
		x,
		&client.DeleteRelationParams{
			SourceIdentifier: 9998,
			TargetIdentifier: 9999,
		},
	)
	assert.True(t, not_found.Is(h))
}
