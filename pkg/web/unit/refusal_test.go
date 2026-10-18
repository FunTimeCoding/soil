package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
	"testing"
)

func TestARefusalCarriesTheReasonFromTheBody(t *testing.T) {
	e := refusal.New(
		"push document",
		answered(http.StatusInternalServerError),
		[]byte(`{"error":"disk full","event_identifier":"alfa"}`),
	)
	assert.True(t, unexpected.Is(e))
	assert.String(t, "push document status: 500: disk full", e.Error())
}

func TestANotFoundRefusalIsNotFound(t *testing.T) {
	e := refusal.New(
		"fetch item",
		answered(http.StatusNotFound),
		[]byte(`{"error":"no such item"}`),
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "fetch item: no such item", e.Error())
}

func TestAnUnavailableRefusalIsUnavailable(t *testing.T) {
	e := refusal.New(
		"hybrid lookup",
		answered(http.StatusServiceUnavailable),
		[]byte(`{"error":"embedder unreachable since 14:02"}`),
	)
	assert.True(t, unavailable.Is(e))
	assert.String(
		t,
		"hybrid lookup: embedder unreachable since 14:02",
		e.Error(),
	)
}

func TestASuccessIsNoRefusal(t *testing.T) {
	refusal.PanicOnFailure("notify", answered(http.StatusNoContent), nil)
	assert.PanicContains(
		t,
		"notify status: 502",
		func() {
			refusal.PanicOnFailure(
				"notify",
				answered(http.StatusBadGateway),
				nil,
			)
		},
	)
}
