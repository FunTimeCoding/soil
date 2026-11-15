package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/reported"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestRefusedPushIsReportedWithItsReason(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(
			`{"error":"disk full","event_identifier":"bravo"}`,
		),
	)
	_, e := client_tester.New(t, s.URL).PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "notes",
			Path:       "charlie.md",
			Body:       "# Charlie\n",
		},
	)
	assert.True(t, reported.Is(e))
	assert.String(t, "push document: disk full (reported as bravo)", e.Error())
}

func TestRefusedListLeavesNoPartialResult(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(
			`{"error":"embedder unreachable since 14:02: no route"}`,
		),
	)
	outcome, e := client_tester.New(t, s.URL).GetList(
		t.Context(),
		&client.GetListParams{Collection: "notes"},
	)
	assert.Nil(t, outcome)
	assert.True(t, unavailable.Is(e))
	assert.String(
		t,
		"list documents: embedder unreachable since 14:02: no route",
		e.Error(),
	)
}
