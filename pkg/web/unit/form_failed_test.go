package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestARedirectForNamesItsFormBesideTheMessage(t *testing.T) {
	w := httptest.NewRecorder()
	form.RedirectFor(
		w,
		httptest.NewRequest(http.MethodPost, "/alfa", nil),
		"/alfa",
		"bravo",
		"charlie",
		url.Values{"delta": {"foxtrot"}},
	)
	assert.Integer(t, http.StatusSeeOther, w.Code)
	assert.String(
		t,
		"/alfa?delta=foxtrot&error=charlie&form=bravo",
		w.Header().Get("Location"),
	)
}

func TestOnlyTheFailedFormSeesTheMessage(t *testing.T) {
	r := httptest.NewRequest(
		http.MethodGet,
		"/alfa?error=charlie&form=bravo",
		nil,
	)
	assert.String(t, "charlie", form.Failed(r, "bravo"))
	assert.String(t, "", form.Failed(r, "delta"))
	assert.String(t, "", form.PageError(r))
}

func TestAPageErrorIsOneNoFormClaims(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/alfa?error=charlie", nil)
	assert.String(t, "charlie", form.PageError(r))
	assert.String(t, "", form.Failed(r, "bravo"))
}
