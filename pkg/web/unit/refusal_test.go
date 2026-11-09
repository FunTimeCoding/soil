package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/reported"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
	"testing"
)

func TestRefusalCarriesReasonFromBody(t *testing.T) {
	e := refusal.New(
		"push document",
		answered(t, http.StatusInternalServerError),
		[]byte(`{"error":"disk full"}`),
	)
	assert.True(t, unexpected.Is(e))
	assert.True(t, classify.Reportable(e))
	assert.String(t, "push document status: 500: disk full", e.Error())
}

func TestCapturedRefusalIsReportedNotCapturedAgain(t *testing.T) {
	e := refusal.New(
		"push document",
		answered(t, http.StatusInternalServerError),
		[]byte(`{"error":"disk full","event_identifier":"alfa"}`),
	)
	assert.True(t, reported.Is(e))
	assert.False(t, classify.Reportable(e))
	assert.String(t, "push document: disk full (reported as alfa)", e.Error())
}

func TestNotFoundRefusalIsNotFound(t *testing.T) {
	e := refusal.New(
		"fetch item",
		answered(t, http.StatusNotFound),
		[]byte(`{"error":"no such item"}`),
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "fetch item: no such item", e.Error())
}

func TestConflictRefusalIsConflict(t *testing.T) {
	e := refusal.New(
		"unmute owner",
		answered(t, http.StatusConflict),
		[]byte(`{"error":"owner not muted: alfa"}`),
	)
	assert.True(t, conflict.Is(e))
	assert.String(t, "unmute owner: owner not muted: alfa", e.Error())
}

func TestUnavailableRefusalIsUnavailable(t *testing.T) {
	e := refusal.New(
		"hybrid lookup",
		answered(t, http.StatusServiceUnavailable),
		[]byte(`{"error":"embedder unreachable since 14:02"}`),
	)
	assert.True(t, unavailable.Is(e))
	assert.String(
		t,
		"hybrid lookup: embedder unreachable since 14:02",
		e.Error(),
	)
}

func TestProxysBadGatewayRefusalIsUnavailable(t *testing.T) {
	e := refusal.New(
		"list items",
		answered(t, http.StatusBadGateway),
		[]byte(`{"error":"no healthy upstream"}`),
	)
	assert.True(t, unavailable.Is(e))
	assert.String(t, "list items: no healthy upstream", e.Error())
}

func TestSuccessIsNoRefusal(t *testing.T) {
	refusal.PanicOnFailure("notify", answered(t, http.StatusNoContent), nil)
	assert.PanicContains(
		t,
		"notify status: 500",
		func() {
			refusal.PanicOnFailure(
				"notify",
				answered(t, http.StatusInternalServerError),
				nil,
			)
		},
	)
}
