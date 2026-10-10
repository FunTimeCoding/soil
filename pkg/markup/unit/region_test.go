package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/markup/constant"
	"github.com/funtimecoding/soil/pkg/markup/region"
	"testing"
)

func TestRegionFindReturnsTheTextBetweenMarkers(t *testing.T) {
	inner, e := region.Find(
		"# Title\n\n<!-- alfa:start -->\n- one\n<!-- alfa:end -->\n\nafter\n",
		constant.FixtureRegion,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "\n- one\n", inner)
}

func TestRegionReplaceKeepsEverythingOutsideTheMarkers(t *testing.T) {
	result, e := region.Replace(
		"before\n<!-- alfa:start -->\nold\n<!-- alfa:end -->\nafter\n",
		constant.FixtureRegion,
		"new",
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"before\n<!-- alfa:start -->\nnew\n<!-- alfa:end -->\nafter\n",
		result,
	)
}

func TestRegionWithoutMarkersIsMissing(t *testing.T) {
	_, e := region.Find("no markers here\n", constant.FixtureRegion)
	assert.ErrorIs(t, e, constant.ErrorRegionMissing)
}

func TestRegionOtherNamesMarkersDoNotCount(t *testing.T) {
	_, e := region.Find(
		"<!-- other:start -->\nx\n<!-- other:end -->\n",
		constant.FixtureRegion,
	)
	assert.ErrorIs(t, e, constant.ErrorRegionMissing)
}

func TestRegionRefusesARepeatedStartMarker(t *testing.T) {
	_, e := region.Find(
		"<!-- alfa:start -->\n<!-- alfa:start -->\n<!-- alfa:end -->\n",
		constant.FixtureRegion,
	)
	assert.Error(t, e)
	assert.String(t, `region "alfa": start marker appears 2 times`, e.Error())
}

func TestRegionRefusesAnEndMarkerWithoutAStart(t *testing.T) {
	_, e := region.Find("x\n<!-- alfa:end -->\n", constant.FixtureRegion)
	assert.Error(t, e)
	assert.String(t, `region "alfa": end marker has no partner`, e.Error())
}

func TestRegionRefusesAStartMarkerWithoutAnEnd(t *testing.T) {
	_, e := region.Replace(
		"<!-- alfa:start -->\nx\n",
		constant.FixtureRegion,
		"new",
	)
	assert.Error(t, e)
	assert.String(t, `region "alfa": start marker has no partner`, e.Error())
}

func TestRegionRefusesAnEndMarkerBeforeTheStart(t *testing.T) {
	_, e := region.Replace(
		"<!-- alfa:end -->\nx\n<!-- alfa:start -->\n",
		constant.FixtureRegion,
		"new",
	)
	assert.Error(t, e)
	assert.String(
		t,
		`region "alfa": end marker comes before start marker`,
		e.Error(),
	)
}
