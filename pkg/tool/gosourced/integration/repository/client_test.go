package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"testing"
)

func TestClientReadsDiff(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, r.Clone, generativeConstant.ModelContextTestToken)
	assert.FatalOnError(t, c.Reachable(t.Context()))
	d, e := c.GetDiff(t.Context(), &client.GetDiffParams{Module: "echo"})
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(d.Entries))
	assert.String(t, "body", d.Entries[0].Class)
	assert.String(t, ". alfa Bravo", d.Entries[0].To.Key)
	assert.String(t, "alfa.go", d.Entries[0].To.File)
	assert.Integer(t, 3, d.Entries[0].To.Line)
	assert.Nil(t, d.Entries[0].Text)
	assert.String(t, "unchanged", d.Counts[0].Class)
	assert.Integer(t, 1, d.Counts[0].Count)
}

func TestClientReadsTextOfChange(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, r.Clone, generativeConstant.ModelContextTestToken)
	d, e := c.GetDiff(
		t.Context(),
		&client.GetDiffParams{Module: "echo", Text: new(true)},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, d.Entries[0].Text)
	assert.Strings(
		t,
		[]string{" func Bravo() int {", "-\treturn 1", "+\treturn 2", " }"},
		*d.Entries[0].Text,
	)
}

func TestClientHearsUnknownModule(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, r.Clone, generativeConstant.ModelContextTestToken)
	_, e := c.GetDiff(t.Context(), &client.GetDiffParams{Module: "golf"})
	assert.True(t, not_found.Is(e))
}

func TestRefusedTokenIsError(t *testing.T) {
	r := diffRepository(t)
	c := mountedClient(t, r.Clone, "hotel-wrong")
	_, e := c.GetDiff(t.Context(), &client.GetDiffParams{Module: "echo"})
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "diff status: 401", e.Error())
	assert.NotNil(t, c.Reachable(t.Context()))
}
