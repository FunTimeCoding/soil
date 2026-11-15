package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/unit/aptly_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestAptlyPublishSignsDistribution(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"Distribution":"rolling","Prefix":".","SourceKind":"local"}`,
		),
	)
	assert.FatalOnError(
		t,
		aptly_tester.New(t, s.URL).UpdatePublish(
			"rolling",
			"/srv/alfa/passphrase",
		),
	)
	r := s.Last()
	assert.String(t, http.MethodPut, r.Method)
	assert.String(t, "/api/publish/:./rolling", r.URL.Path)
	assert.StringContains(
		t,
		`"PassphraseFile":"/srv/alfa/passphrase"`,
		s.LastBody(),
	)
}

func TestAptlyPublishRefusesWhatAptlyRefused(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(
			`{"error":"unable to update: unable to sign: gpg: signing failed"}`,
		),
	)
	assert.NotNil(
		t,
		aptly_tester.New(t, s.URL).UpdatePublish(
			"rolling",
			"/srv/alfa/passphrase",
		),
	)
}
