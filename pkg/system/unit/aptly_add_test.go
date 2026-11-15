package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/unit/aptly_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestAptlyAddAcceptsWhatAptlyAdded(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"FailedFiles":[],"Report":{"Warnings":[],"Added":["alfa_1.0.0_amd64 added"],"Removed":[]}}`,
		),
	)
	assert.FatalOnError(
		t,
		aptly_tester.New(t, s.URL).AddToRepository(
			"rolling",
			"alfa-1.0.0",
			false,
		),
	)
	assert.String(
		t,
		"",
		s.Last().URL.Query().Get(constant.AptlyForceReplaceParameter),
	)
}

func TestAptlyAddRefusesWhatAptlyFailed(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"FailedFiles":["/srv/upload/alfa-1.0.0/alfa_1.0.0-1_amd64.deb"],"Report":{"Warnings":["Unable to add package: package already exists and is different: alfa_1.0.0_amd64"],"Added":[],"Removed":[]}}`,
		),
	)
	e := aptly_tester.New(t, s.URL).AddToRepository(
		"rolling",
		"alfa-1.0.0",
		false,
	)
	assert.True(t, conflict.Is(e))
	assert.StringContains(t, "already exists and is different", e.Error())
}

func TestAptlyAddForcesReplacement(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"FailedFiles":[],"Report":{"Warnings":[],"Added":["alfa_1.0.0_amd64 added"],"Removed":["alfa_1.0.0_amd64 removed"]}}`,
		),
	)
	assert.FatalOnError(
		t,
		aptly_tester.New(t, s.URL).AddToRepository(
			"rolling",
			"alfa-1.0.0",
			true,
		),
	)
	assert.String(
		t,
		"1",
		s.Last().URL.Query().Get(constant.AptlyForceReplaceParameter),
	)
}
