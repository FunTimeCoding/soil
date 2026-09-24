package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestCoverageEndToEnd(t *testing.T) {
	root := module(t)
	profile := coverage.Run(root, "", constant.AllPackages)
	assert.True(t, system.FileExists(profile))
	r := crap.Build(
		index.Load(root, constant.AllPackages),
		coverage.Functions(root, profile),
		constant.Pessimistic,
	)
	assert.Count(t, 3, r.Entries)
	byName := map[string]float64{}
	missing := map[string]bool{}

	for _, e := range r.Entries {
		byName[e.Function.Name] = e.Coverage
		missing[e.Function.Name] = e.Missing
	}

	assert.Float(t, 100, byName["Covered"])
	assert.True(t, byName["Half"] > 0 && byName["Half"] < 100)
	assert.Float(t, 0, byName["Untested"])
	assert.False(t, missing["Untested"])
	assert.Float(t, 12, r.Sorted()[0].Score)
}
