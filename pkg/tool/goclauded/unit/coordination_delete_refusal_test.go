package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"net/http"
	"testing"
)

func TestDeletingSessionWithTranscriptNeedsConfirmation(t *testing.T) {
	s := base.NewWithSessionFile(t, "doomed", "some-slug")
	r, e := s.RESTClient(t).DeleteSessionByIdentifierWithResponse(
		t.Context(),
		"doomed",
		&client.DeleteSessionByIdentifierParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusConflict, r.StatusCode())
	assert.String(
		t,
		"session has a transcript, confirmation required",
		r.JSON409.Error,
	)
	assert.True(t, s.Store.GetSession("doomed") != nil)
}
