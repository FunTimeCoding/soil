package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/mdstat"
	"testing"
)

func TestParseHealthyMirror(t *testing.T) {
	a := mdstat.Parse(fixture.Read("disk", "mdstat_healthy.txt"))
	assert.Count(t, 1, a)
	assert.String(t, "md1", a[0].Name)
	assert.String(t, "active", a[0].State)
	assert.String(t, "raid1", a[0].Level)
	assert.Integer(t, 2, a[0].Required)
	assert.Integer(t, 2, a[0].Active)
	assert.Integer(t, 0, a[0].Failed)
	assert.False(t, a[0].Rebuilding)
}

func TestParseDegradedMirrorRecovering(t *testing.T) {
	a := mdstat.Parse(fixture.Read("disk", "mdstat_degraded_derived.txt"))
	assert.Count(t, 1, a)
	assert.Integer(t, 2, a[0].Required)
	assert.Integer(t, 1, a[0].Active)
	assert.Integer(t, 1, a[0].Failed)
	assert.True(t, a[0].Rebuilding)
}

func TestParseKeepsArraysApart(t *testing.T) {
	a := mdstat.Parse(fixture.Read("disk", "mdstat_two_arrays_derived.txt"))
	assert.Count(t, 2, a)
	assert.Integer(t, 2, a[0].Active)
	assert.Integer(t, 0, a[0].Failed)
	assert.String(t, "md1", a[1].Name)
	assert.String(t, "raid1", a[1].Level)
	assert.Integer(t, 1, a[1].Active)
	assert.Integer(t, 1, a[1].Failed)
}

func TestParseWithoutArraysIsEmpty(t *testing.T) {
	assert.Count(
		t,
		0,
		mdstat.Parse("Personalities : \nunused devices: <none>\n"),
	)
}
