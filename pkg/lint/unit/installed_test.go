package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/installed"
	"github.com/funtimecoding/soil/pkg/stamp"
	"os"
	"testing"
)

func TestReadTakesStampFromBuildInformation(t *testing.T) {
	path, e := os.Executable()
	assert.FatalOnError(t, e)
	b, okay := installed.Read(path)
	assert.True(t, okay)
	assert.String(t, "github.com/funtimecoding/soil", b.Module)
	assert.String(t, stamp.New().Version, b.Version)
}

func TestBinaryBelowLatestTagIsStale(t *testing.T) {
	assert.True(t, installed.Stale("v0.11.169", "v0.11.170"))
}

func TestBinaryAtLatestTagIsNotStale(t *testing.T) {
	assert.False(t, installed.Stale("v0.11.170", "v0.11.170"))
}

func TestDirtyBinaryAtLatestTagIsNotStale(t *testing.T) {
	assert.False(t, installed.Stale("v0.11.170+dirty", "v0.11.170"))
}

func TestBinaryBuiltAfterLatestTagIsNotStale(t *testing.T) {
	assert.False(
		t,
		installed.Stale(
			"v0.11.171-0.20261004120000-fd20dc25a022+dirty",
			"v0.11.170",
		),
	)
}

func TestBinaryBuiltBeforeLatestTagIsStale(t *testing.T) {
	assert.True(
		t,
		installed.Stale("v0.11.170-0.20261003120000-f5dce445d000", "v0.11.170"),
	)
}

func TestRepositoryWithoutTagsLeavesNothingStale(t *testing.T) {
	assert.False(t, installed.Stale("v0.11.170", ""))
}
