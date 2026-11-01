package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func TestAptlyPublishSignsTheDistribution(t *testing.T) {
	var seen aptlySeen
	c := fakeAptly(
		t,
		http.StatusOK,
		`{"Distribution":"rolling","Prefix":".","SourceKind":"local"}`,
		&seen,
	)
	assert.FatalOnError(t, c.UpdatePublish("rolling", "/srv/alfa/passphrase"))
	assert.String(t, http.MethodPut, seen.Method)
	assert.String(t, "/api/publish/:./rolling", seen.Path)
	assert.StringContains(
		t,
		`"PassphraseFile":"/srv/alfa/passphrase"`,
		seen.Body,
	)
}

func TestAptlyPublishRefusesWhatAptlyRefused(t *testing.T) {
	var seen aptlySeen
	c := fakeAptly(
		t,
		http.StatusInternalServerError,
		`{"error":"unable to update: unable to sign: gpg: signing failed"}`,
		&seen,
	)
	assert.NotNil(t, c.UpdatePublish("rolling", "/srv/alfa/passphrase"))
}
