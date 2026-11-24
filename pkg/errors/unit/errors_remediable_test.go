package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/remediable"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"net/http"
	"testing"
)

func TestRemediableWrapKeepsItsCause(t *testing.T) {
	e := remediable.Wrap(
		timeout.Format("browser tab: did not answer"),
		"wake_tab",
		"bring the tab forward",
	)
	assert.String(
		t,
		"bring the tab forward: browser tab: did not answer",
		e.Error(),
	)
	assert.String(t, "wake_tab", remediable.Remedy(e))
	assert.True(t, remediable.Is(e))
	assert.True(t, timeout.Is(e))
}

func TestRemediableWinsOverItsCause(t *testing.T) {
	e := remediable.Wrap(
		timeout.Format("browser tab: did not answer"),
		"wake_tab",
		"bring the tab forward",
	)
	assert.String(t, "remediable", classify.Class(e))
	assert.Integer(t, http.StatusConflict, classify.Answer(e).Status)
	assert.False(t, classify.Reportable(e))
}

func TestRemediableFormatHasNoCause(t *testing.T) {
	e := remediable.Format("authorize", "authorize %s again", "alfa")
	assert.String(t, "authorize alfa again", e.Error())
	assert.String(t, "authorize", remediable.Remedy(e))
	assert.Nil(t, errors.Unwrap(e))
}

func TestRemedyOfPlainErrorIsEmpty(t *testing.T) {
	assert.String(t, "", remediable.Remedy(errors.New("disk full")))
	assert.False(t, remediable.Is(errors.New("disk full")))
}
