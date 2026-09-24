package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"testing"
)

func TestBuildPolicies(t *testing.T) {
	i := index.New("/r")
	i.Add(function.New("m", "/r/a.go", 1, "A", 4))
	i.Add(function.New("m", "/r/b.go", 1, "B", 4))
	coverage := map[string]float64{"m/a.go:1": 50}
	pessimistic := crap.Build(i, coverage, constant.Pessimistic)
	assert.Count(t, 2, pessimistic.Entries)
	assert.Float(t, 6, pessimistic.Entries[0].Score)
	assert.Float(t, 20, pessimistic.Entries[1].Score)
	assert.True(t, pessimistic.Entries[1].Missing)
	optimistic := crap.Build(i, coverage, constant.Optimistic)
	assert.Float(t, 4, optimistic.Entries[1].Score)
	skip := crap.Build(i, coverage, constant.Skip)
	assert.Count(t, 1, skip.Entries)
	assert.Integer(t, 1, skip.Skipped)
}

func TestReportOrderAndSummary(t *testing.T) {
	i := index.New("/r")
	i.Add(function.New("m", "/r/a.go", 1, "A", 2))
	i.Add(function.New("m", "/r/b.go", 1, "B", 6))
	r := crap.Build(
		i,
		map[string]float64{"m/a.go:1": 100, "m/b.go:1": 0},
		constant.Pessimistic,
	)
	sorted := r.Sorted()
	assert.String(t, "B", sorted[0].Function.Name)
	assert.Float(t, 42, sorted[0].Score)
	assert.Count(t, 1, r.Above(30))
	assert.Float(t, 44, r.Combined())
	assert.Float(t, 22, r.Average())
	assert.String(t, "b.go:1", sorted[0].Location("/r"))
}
