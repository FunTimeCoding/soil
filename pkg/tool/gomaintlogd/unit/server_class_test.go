package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	"net/http"
	"testing"
)

func TestDeleteOfUnknownEntryIsAnsweredNotFound(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().DeleteEntryWithResponse(t.Context(), 9999)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.FatalNotNil(t, r.JSON404)
	assert.String(t, "entry not found: 9999", r.JSON404.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestEntryWithoutActionIsAnsweredBadRequest(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().PostEntryWithResponse(
		t.Context(),
		client.PostEntryJSONRequestBody{User: "alice"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "action is required", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestStatusFaultIsCaptured(t *testing.T) {
	s := base.New(t)
	s.Execute("ALTER TABLE maintenance_log_entries RENAME TO entries_gone")
	r, e := s.GeneratedClient().GetStatusWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.FatalNotNil(t, r.JSON500)
	assert.Count(t, 1, s.Reporter.Events())
}
