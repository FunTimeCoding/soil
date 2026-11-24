package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/installed"
	"github.com/funtimecoding/soil/pkg/lint/unit/installed_tester"
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

func TestCommandInDownstreamCheckoutBelongsToDownstream(t *testing.T) {
	o := installed_tester.New(t)
	upstream := o.Checkout("upstream", "example.com/upstream", nil, "alfa")
	o.Checkout(
		"downstream",
		"example.com/downstream",
		[]string{"example.com/upstream"},
		"alfa",
	)
	assert.String(
		t,
		"example.com/downstream",
		installed.OwnerOf(upstream, "alfa"),
	)
}

func TestCommandInOneCheckoutBelongsToIt(t *testing.T) {
	o := installed_tester.New(t)
	upstream := o.Checkout("upstream", "example.com/upstream", nil, "bravo")
	o.Checkout(
		"downstream",
		"example.com/downstream",
		[]string{"example.com/upstream"},
	)
	assert.String(
		t,
		"example.com/upstream",
		installed.OwnerOf(upstream, "bravo"),
	)
}

func TestCommandRepackagedDownstreamBelongsToUpstream(t *testing.T) {
	o := installed_tester.New(t)
	upstream := o.Checkout("upstream", "example.com/upstream", nil, "delta")
	downstream := o.Checkout(
		"downstream",
		"example.com/downstream",
		[]string{"example.com/upstream"},
		"delta",
	)
	o.Main(downstream, "delta", "example.com/upstream/pkg/tool/delta")
	assert.String(
		t,
		"example.com/upstream",
		installed.OwnerOf(upstream, "delta"),
	)
}

func TestCommandInUnrelatedCheckoutsBelongsToNone(t *testing.T) {
	o := installed_tester.New(t)
	first := o.Checkout("first", "example.com/first", nil, "charlie")
	o.Checkout("second", "example.com/second", nil, "charlie")
	assert.String(t, "", installed.OwnerOf(first, "charlie"))
}
