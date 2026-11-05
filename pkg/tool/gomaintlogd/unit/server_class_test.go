package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/web_service_tester"
	"net/http"
	"testing"
)

func TestDeleteOfUnknownEntryIsAnsweredNotFound(t *testing.T) {
	o := web_service_tester.New(t)
	r, e := o.Client.DeleteEntryWithResponse(t.Context(), 9999)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.FatalNotNil(t, r.JSON404)
	assert.String(t, "entry not found: 9999", r.JSON404.Error)
	assert.Count(t, 0, o.Reporter.Events())
}

func TestEntryWithoutActionIsAnsweredBadRequest(t *testing.T) {
	o := web_service_tester.New(t)
	r, e := o.Client.PostEntryWithResponse(
		t.Context(),
		client.PostEntryJSONRequestBody{User: "alice"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "action is required", r.JSON400.Error)
	assert.Count(t, 0, o.Reporter.Events())
}

func TestStatusFaultIsCaptured(t *testing.T) {
	o := web_service_tester.New(t)
	o.Execute("ALTER TABLE maintenance_log_entries RENAME TO entries_gone")
	r, e := o.Client.GetStatusWithResponse(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, r.StatusCode())
	assert.FatalNotNil(t, r.JSON500)
	assert.Count(t, 1, o.Reporter.Events())
}
