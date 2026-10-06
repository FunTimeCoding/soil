package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/build"
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

func TestAMainFileBuildsItsPackage(t *testing.T) {
	assert.String(t, "./cmd/golint", build.Package("cmd/golint/main.go"))
}

func TestTagsCarryTimeZonesWithoutRequest(t *testing.T) {
	assert.String(t, "timetzdata", build.Tags(""))
}

func TestTagsKeepTimeZonesBesideRequestedTags(t *testing.T) {
	assert.String(t, "timetzdata,netgo", build.Tags("netgo"))
}
