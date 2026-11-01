package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net/http"
	"testing"
)

func TestAptlyAddAcceptsWhatAptlyAdded(t *testing.T) {
	var seen aptlySeen
	c := fakeAptly(
		t,
		http.StatusOK,
		`{"FailedFiles":[],"Report":{"Warnings":[],"Added":["alfa_1.0.0_amd64 added"],"Removed":[]}}`,
		&seen,
	)
	assert.FatalOnError(t, c.AddToRepository("rolling", "alfa-1.0.0", false))
	assert.String(t, "", seen.Query.Get(constant.AptlyForceReplaceParameter))
}

func TestAptlyAddRefusesWhatAptlyFailed(t *testing.T) {
	var seen aptlySeen
	c := fakeAptly(
		t,
		http.StatusOK,
		`{"FailedFiles":["/srv/upload/alfa-1.0.0/alfa_1.0.0-1_amd64.deb"],"Report":{"Warnings":["Unable to add package: package already exists and is different: alfa_1.0.0_amd64"],"Added":[],"Removed":[]}}`,
		&seen,
	)
	e := c.AddToRepository("rolling", "alfa-1.0.0", false)
	assert.True(t, conflict.Is(e))
	assert.StringContains(t, "already exists and is different", e.Error())
}

func TestAptlyAddForcesTheReplacement(t *testing.T) {
	var seen aptlySeen
	c := fakeAptly(
		t,
		http.StatusOK,
		`{"FailedFiles":[],"Report":{"Warnings":[],"Added":["alfa_1.0.0_amd64 added"],"Removed":["alfa_1.0.0_amd64 removed"]}}`,
		&seen,
	)
	assert.FatalOnError(t, c.AddToRepository("rolling", "alfa-1.0.0", true))
	assert.String(t, "1", seen.Query.Get(constant.AptlyForceReplaceParameter))
}
