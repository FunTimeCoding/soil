package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/generation"
	"github.com/funtimecoding/soil/pkg/source/unit/generation_tester"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerationOpenRemovesSiblingUntouchedPastRetention(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	generation_tester.Aged(t, root, "alfa", now.Add(-4*24*time.Hour))
	generation.Open(root, "bravo", now, generation_tester.EveryName)
	assert.Any(t, []string{"bravo"}, system.DirectoryContent(root))
}

func TestGenerationOpenKeepsSiblingTouchedWithinRetention(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	generation_tester.Aged(t, root, "alfa", now.Add(-2*24*time.Hour))
	generation.Open(root, "bravo", now, generation_tester.EveryName)
	assert.Any(t, []string{"alfa", "bravo"}, system.DirectoryContent(root))
}

func TestGenerationOpenRefreshesItsOwnMarker(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	generation_tester.Aged(t, root, "alfa", now.Add(-2*24*time.Hour))
	generation.Open(root, "alfa", now, generation_tester.EveryName)
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
	generation.Open(root, "bravo", time.Now(), generation_tester.EveryName)
	assert.Any(t, []string{"bravo"}, system.DirectoryContent(root))
}

func TestGenerationPruneRemovesEverySiblingButKeptOne(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	generation_tester.Aged(t, root, "alfa", now)
	generation_tester.Aged(t, root, "charlie", now)
	generation_tester.Aged(t, root, "bravo", now)
	assert.Any(
		t,
		[]string{"alfa", "charlie"},
		generation.Prune(root, "bravo", generation_tester.EveryName),
	)
	assert.Any(t, []string{"bravo"}, system.DirectoryContent(root))
}

func TestGenerationOpenLeavesSiblingItDoesNotOwn(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	generation_tester.Aged(t, root, "alfa", now.Add(-4*24*time.Hour))
	generation.Open(root, "bravo", now, func(string) bool { return false })
	assert.Any(t, []string{"alfa", "bravo"}, system.DirectoryContent(root))
}
