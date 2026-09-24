package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gohook/job"
	"testing"
)

func TestMatchesPrefix(t *testing.T) {
	j := job.New("task lint-manifest", "strata/manifest/")
	assert.True(t, j.Matches([]string{"strata/manifest/deploy.yaml"}))
	assert.False(t, j.Matches([]string{"strata/manifests/deploy.yaml"}))
	assert.False(t, j.Matches([]string{"pkg/strata/manifest/x.yaml"}))
}

func TestMatchesGlob(t *testing.T) {
	j := job.New("task lint", "**/*.go")
	assert.True(t, j.Matches([]string{"pkg/alfa/bravo.go"}))
	assert.True(t, j.Matches([]string{"main.go"}))
	assert.False(t, j.Matches([]string{"pkg/alfa/bravo.yaml"}))
	assert.True(t, job.New("x", "*.md").Matches([]string{"README.md"}))
	assert.False(t, job.New("x", "*.md").Matches([]string{"doc/README.md"}))
}

func TestMatchesExact(t *testing.T) {
	j := job.New("go mod tidy", "go.mod")
	assert.True(t, j.Matches([]string{"go.mod"}))
	assert.False(t, j.Matches([]string{"pkg/go.mod"}))
}

func TestMatchesAnyPattern(t *testing.T) {
	j := job.New("x", "go.mod", "**/*.go")
	assert.True(t, j.Matches([]string{"pkg/a.go"}))
	assert.True(t, j.Matches([]string{"go.mod"}))
	assert.False(t, j.Matches([]string{"README.md"}))
}

func TestMatchesWithoutPaths(t *testing.T) {
	j := job.New("task always")
	assert.True(t, j.Matches(nil))
	assert.True(t, j.Matches([]string{"anything"}))
}

func TestMatchesNoChanges(t *testing.T) {
	assert.False(t, job.New("x", "**/*.go").Matches(nil))
}
