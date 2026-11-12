package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/generation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerationOpenRemovesSiblingUntouchedPastRetention(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	agedGeneration(t, root, "alfa", now.Add(-4*24*time.Hour))
	generation.Open(root, "bravo", now, everyName)
	assert.Any(t, []string{"bravo"}, entryNames(t, root))
}

func TestGenerationOpenKeepsSiblingTouchedWithinRetention(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	agedGeneration(t, root, "alfa", now.Add(-2*24*time.Hour))
	generation.Open(root, "bravo", now, everyName)
	assert.Any(t, []string{"alfa", "bravo"}, entryNames(t, root))
}

func TestGenerationOpenRefreshesItsOwnMarker(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	agedGeneration(t, root, "alfa", now.Add(-2*24*time.Hour))
	generation.Open(root, "alfa", now, everyName)
	i, e := os.Stat(filepath.Join(root, "alfa", constant.GenerationMarker))
	assert.FatalOnError(t, e)
	assert.True(t, now.Sub(i.ModTime()) < time.Minute)
}

func TestGenerationOpenRemovesLeftoverTrash(t *testing.T) {
	root := t.TempDir()
	assert.FatalOnError(
		t,
		os.MkdirAll(
			filepath.Join(root, join.Empty(constant.GenerationTrash, "charlie")),
			0755,
		),
	)
	generation.Open(root, "bravo", time.Now(), everyName)
	assert.Any(t, []string{"bravo"}, entryNames(t, root))
}

func TestGenerationPruneRemovesEverySiblingButKeptOne(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	agedGeneration(t, root, "alfa", now)
	agedGeneration(t, root, "charlie", now)
	agedGeneration(t, root, "bravo", now)
	assert.Any(
		t,
		[]string{"alfa", "charlie"},
		generation.Prune(root, "bravo", everyName),
	)
	assert.Any(t, []string{"bravo"}, entryNames(t, root))
}

func TestGenerationOpenLeavesSiblingItDoesNotOwn(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	agedGeneration(t, root, "alfa", now.Add(-4*24*time.Hour))
	generation.Open(root, "bravo", now, func(string) bool { return false })
	assert.Any(t, []string{"alfa", "bravo"}, entryNames(t, root))
}
