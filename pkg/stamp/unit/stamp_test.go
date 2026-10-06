package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"runtime/debug"
	"testing"
)

func TestReadTakesEverythingFromTheBuildInformation(t *testing.T) {
	assert.Any(
		t,
		&stamp.Stamp{
			Version:    "v1.2.3",
			GitHash:    "fd20dc2",
			CommitDate: "2026-10-03T21:38:13Z",
			Module:     "github.com/funtimecoding/soil",
			Dirty:      true,
		},
		stamp.Read(
			buildInformation(
				"v1.2.3",
				debug.BuildSetting{
					Key:   "vcs.revision",
					Value: "fd20dc25a0228d19ab7d984e3946ef16ed65cb15",
				},
				debug.BuildSetting{
					Key:   "vcs.time",
					Value: "2026-10-03T21:38:13Z",
				},
				debug.BuildSetting{Key: "vcs.modified", Value: "true"},
			),
		),
	)
}

func TestADevelopmentBuildFallsBackToDefaults(t *testing.T) {
	assert.Any(
		t,
		&stamp.Stamp{
			Version:    "empty",
			GitHash:    "empty",
			CommitDate: "empty",
			Module:     "github.com/funtimecoding/soil",
		},
		stamp.Read(buildInformation(constant.DevelopmentVersion)),
	)
}

func TestNoBuildInformationFallsBackToDefaults(t *testing.T) {
	assert.Any(
		t,
		&stamp.Stamp{
			Version:    "empty",
			GitHash:    "empty",
			CommitDate: "empty",
		},
		stamp.Read(nil),
	)
}

func TestATaggedVersionDisplaysAsItIs(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(buildInformation("v0.11.170")).DisplayVersion(),
	)
}

func TestADirtyTaggedVersionDisplaysWithoutTheSuffix(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(buildInformation("v0.11.170+dirty")).DisplayVersion(),
	)
}

func TestAPseudoVersionDisplaysTheTagItFollows(t *testing.T) {
	assert.String(
		t,
		"v0.11.170 (untagged)",
		stamp.Read(
			buildInformation("v0.11.171-0.20261003213813-fd20dc25a022+dirty"),
		).DisplayVersion(),
	)
}

func TestAPseudoVersionWithoutAnyTagDisplaysAsUntagged(t *testing.T) {
	assert.String(
		t,
		"untagged",
		stamp.Read(
			buildInformation("v0.0.0-20261003213813-fd20dc25a022"),
		).DisplayVersion(),
	)
}

func TestADirtyTaggedBuildReleasesAsItsTag(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(buildInformation("v0.11.170+dirty")).Tag(),
	)
}

func TestAnUntaggedBuildReleasesAsTheTagItFollows(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(
			buildInformation("v0.11.171-0.20261003213813-fd20dc25a022+dirty"),
		).Tag(),
	)
}

func TestABuildBeforeAnyTagHasNoTag(t *testing.T) {
	assert.String(
		t,
		"",
		stamp.Read(
			buildInformation("v0.0.0-20261003213813-fd20dc25a022"),
		).Tag(),
	)
}

func TestADevelopmentBuildHasNoTag(t *testing.T) {
	assert.String(
		t,
		"",
		stamp.Read(buildInformation(constant.DevelopmentVersion)).Tag(),
	)
}

func TestAPseudoVersionWithoutRevisionCarriesItsCommit(t *testing.T) {
	assert.String(
		t,
		"fd20dc2",
		stamp.Read(
			buildInformation("v0.11.171-0.20261003213813-fd20dc25a022"),
		).GitHash,
	)
}

func TestATaggedVersionWithoutRevisionHasNoCommit(t *testing.T) {
	assert.String(t, "empty", stamp.Read(buildInformation("v0.11.170")).GitHash)
}
