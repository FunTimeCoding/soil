package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/module"
	"testing"
)

func TestModuleTakesTheLatestReleaseFromTheUpdate(t *testing.T) {
	m := modules(
		`{"Path": "a", "Version": "v1.0.0", "Time": "2020-01-04T00:00:00Z", "Update": {"Path": "a", "Version": "v1.1.0", "Time": "2026-09-01T00:00:00Z"}}`,
	)[0]
	assert.String(t, "v1.0.0", m.Version)
	assert.String(t, "v1.1.0", m.Latest)
	assert.Integer(t, 1, m.Months)
	assert.False(t, m.HasConcerns())
}

func TestModuleIsStaleAtTheThreshold(t *testing.T) {
	m := modules(
		`{"Path": "a", "Version": "v1.0.0", "Time": "2025-10-11T00:00:00Z"}`,
	)[0]
	assert.Integer(t, 12, m.Months)
	assert.True(t, m.HasConcern(constant.ReleaseStale))
}

func TestModuleIsFreshBelowTheThreshold(t *testing.T) {
	m := modules(
		`{"Path": "a", "Version": "v1.0.0", "Time": "2025-10-12T00:00:00Z"}`,
	)[0]
	assert.Integer(t, 11, m.Months)
	assert.False(t, m.HasConcerns())
}

func TestModuleWithoutReleaseTimeHasNoConcern(t *testing.T) {
	assert.False(
		t,
		modules(`{"Path": "a", "Version": "v1.0.0"}`)[0].HasConcerns(),
	)
}

func TestModulePseudoVersionIsUntagged(t *testing.T) {
	assert.True(
		t,
		modules(
			`{"Path": "a", "Version": "v1.5.4-0.20250319132907-e064f32e3674", "Time": "2025-03-19T13:29:07Z"}`,
		)[0].Untagged,
	)
}

func TestSortStalestPutsModulesWithoutReleaseTimeLast(t *testing.T) {
	v := modules(
		`{"Path": "none", "Version": "v1.0.0"}
{"Path": "new", "Version": "v1.0.0", "Time": "2026-01-01T00:00:00Z"}
{"Path": "old", "Version": "v1.0.0", "Time": "2019-04-18T00:00:00Z"}`,
	)
	module.SortStalest(v)
	assert.String(t, "old", v[0].Path)
	assert.String(t, "new", v[1].Path)
	assert.String(t, "none", v[2].Path)
}
