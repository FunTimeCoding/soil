package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"net/http"
	"testing"
)

func TestARefusedPushIsUnexpectedWithItsReason(t *testing.T) {
	_, e := refusedClient(
		t,
		http.StatusInternalServerError,
		`{"error":"disk full","event_identifier":"bravo"}`,
	).PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "notes",
			Path:       "charlie.md",
			Body:       "# Charlie\n",
		},
	)
	assert.True(t, unexpected.Is(e))
	assert.String(t, "push document status: 500: disk full", e.Error())
}

func TestARefusedListLeavesNoPartialResult(t *testing.T) {
	outcome, e := refusedClient(
		t,
		http.StatusServiceUnavailable,
		`{"error":"embedder unreachable since 14:02: no route"}`,
	).GetList(t.Context(), &client.GetListParams{Collection: "notes"})
	assert.Nil(t, outcome)
	assert.True(t, unavailable.Is(e))
	assert.String(
		t,
		"list documents: embedder unreachable since 14:02: no route",
		e.Error(),
	)
}
