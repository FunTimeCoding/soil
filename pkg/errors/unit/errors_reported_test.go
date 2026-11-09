package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/reported"
	"net/http"
	"testing"
)

func TestReportedClassifiesByType(t *testing.T) {
	e := reported.New("retarget references", "disk full", "alfa")
	assert.String(
		t,
		"retarget references: disk full (reported as alfa)",
		e.Error(),
	)
	assert.String(t, "alfa", e.Event)
	assert.True(t, reported.Is(e))
	assert.False(t, reported.Is(errors.New("disk full")))
}

func TestReportedFailureAnswersBadGatewayUncaptured(t *testing.T) {
	e := reported.New("retarget references", "disk full", "alfa")
	assert.String(t, "reported", classify.Class(e))
	assert.Integer(t, http.StatusBadGateway, classify.Answer(e).Status)
	assert.False(t, classify.Reportable(e))
}
