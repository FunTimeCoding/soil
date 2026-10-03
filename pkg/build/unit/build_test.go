package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func TestParseOriginReadsTheCommitAndTime(t *testing.T) {
	v := build.ParseOrigin(originOutput())
	assert.String(t, "72ce25dbd0b797a65759efe648f5838cc9709481", v.Hash)
	assert.String(t, "2026-09-27T14:17:43Z", v.Time)
}

func TestParseOriginWithoutOriginReadsNoCommit(t *testing.T) {
	v := build.ParseOrigin(
		`{"Path":"alfa","Version":"v1.0.0","Time":"2026-01-01T00:00:00Z"}`,
	)
	assert.String(t, "", v.Hash)
	assert.String(t, "2026-01-01T00:00:00Z", v.Time)
}

func TestParseOriginRefusesGarbage(t *testing.T) {
	assert.Nil(t, build.ParseOrigin("not json"))
}

func TestLinkerFlagsSetEveryVariable(t *testing.T) {
	assert.String(
		t,
		join.Empty(
			"-X main.Version=v1.2.3 -X main.GitHash=abc1234 ",
			"-X main.BuildDate=2026-01-01T00:00:00Z ",
			"-X github.com/funtimecoding/soil/pkg/stamp/constant.Module=example.com/tool ",
			"-X github.com/funtimecoding/soil/pkg/stamp/constant.Dirty=1",
		),
		build.LinkerFlags(
			"v1.2.3",
			"abc1234",
			"2026-01-01T00:00:00Z",
			"example.com/tool",
			true,
		),
	)
}

func TestShortHashTruncates(t *testing.T) {
	assert.String(
		t,
		"72ce25d",
		build.ShortHash("72ce25dbd0b797a65759efe648f5838cc9709481"),
	)
}

func TestShortHashKeepsAShortValue(t *testing.T) {
	assert.String(t, "abc", build.ShortHash("abc"))
}

func TestTagsCarryTimeZonesWithoutRequest(t *testing.T) {
	assert.String(t, "timetzdata", build.Tags(""))
}

func TestTagsKeepTimeZonesBesideRequestedTags(t *testing.T) {
	assert.String(t, "timetzdata,netgo", build.Tags("netgo"))
}

func TestBuild(t *testing.T) {
	d := git.FindDirectory()

	if d == "" {
		t.Skip("no git directory")
	}

	assert.True(t, len(build.Date()) > 0)
	assert.True(t, len(build.GitHash()) > 0)

	if len(git.Tags(d)) == 0 {
		t.Skip("no git tags")
	}

	assert.True(t, len(build.GitTag()) > 0)
}
