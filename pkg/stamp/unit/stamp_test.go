package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"github.com/funtimecoding/soil/pkg/stamp/unit/stamp_tester"
	"runtime/debug"
	"testing"
)

func TestReadTakesEverythingFromBuildInformation(t *testing.T) {
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
			stamp_tester.BuildInformation(
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

func TestDevelopmentBuildFallsBackToDefaults(t *testing.T) {
	assert.Any(
		t,
		&stamp.Stamp{
			Version:    "empty",
			GitHash:    "empty",
			CommitDate: "empty",
			Module:     "github.com/funtimecoding/soil",
		},
		stamp.Read(stamp_tester.BuildInformation(constant.DevelopmentVersion)),
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

func TestTaggedVersionDisplaysAsItIs(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(stamp_tester.BuildInformation("v0.11.170")).DisplayVersion(),
	)
}

func TestDirtyTaggedVersionDisplaysWithoutSuffix(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(stamp_tester.BuildInformation("v0.11.170+dirty")).DisplayVersion(),
	)
}

func TestPseudoVersionDisplaysTagItFollows(t *testing.T) {
	assert.String(
		t,
		"v0.11.170 (untagged)",
		stamp.Read(
			stamp_tester.BuildInformation(
				"v0.11.171-0.20261003213813-fd20dc25a022+dirty",
			),
		).DisplayVersion(),
	)
}

func TestPseudoVersionWithoutAnyTagDisplaysAsUntagged(t *testing.T) {
	assert.String(
		t,
		"untagged",
		stamp.Read(
			stamp_tester.BuildInformation("v0.0.0-20261003213813-fd20dc25a022"),
		).DisplayVersion(),
	)
}

func TestDirtyTaggedBuildReleasesAsItsTag(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(stamp_tester.BuildInformation("v0.11.170+dirty")).Tag(),
	)
}

func TestUntaggedBuildReleasesAsTagItFollows(t *testing.T) {
	assert.String(
		t,
		"v0.11.170",
		stamp.Read(
			stamp_tester.BuildInformation(
				"v0.11.171-0.20261003213813-fd20dc25a022+dirty",
			),
		).Tag(),
	)
}

func TestBuildBeforeAnyTagHasNoTag(t *testing.T) {
	assert.String(
		t,
		"",
		stamp.Read(
			stamp_tester.BuildInformation("v0.0.0-20261003213813-fd20dc25a022"),
		).Tag(),
	)
}

func TestDevelopmentBuildHasNoTag(t *testing.T) {
	assert.String(
		t,
		"",
		stamp.Read(stamp_tester.BuildInformation(constant.DevelopmentVersion)).Tag(),
	)
}

func TestPseudoVersionWithoutRevisionCarriesItsCommit(t *testing.T) {
	assert.String(
		t,
		"fd20dc2",
		stamp.Read(
			stamp_tester.BuildInformation(
				"v0.11.171-0.20261003213813-fd20dc25a022",
			),
		).GitHash,
	)
}

func TestTaggedVersionWithoutRevisionHasNoCommit(t *testing.T) {
	assert.String(
		t,
		"empty",
		stamp.Read(stamp_tester.BuildInformation("v0.11.170")).GitHash,
	)
}
