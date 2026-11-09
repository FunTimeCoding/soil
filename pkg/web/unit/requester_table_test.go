package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
	"testing"
)

func TestTableAnswersGatewayStatusesUnavailable(t *testing.T) {
	for _, status := range []int{
		http.StatusTooManyRequests,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	} {
		s, _ := newScriptedServer(t, "", status)
		_, e := newRequester(t, s.URL).Bytes(request.Get("/items"))
		assert.True(t, unavailable.Is(e))
	}
}

func TestTableAnswersConflictAsConflict(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusConflict)
	_, e := newRequester(t, s.URL).Bytes(request.Get("/items"))
	assert.True(t, conflict.Is(e))
}

func TestTableKeepsOurOwnRefusalsUnexpected(t *testing.T) {
	for _, status := range []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusUnprocessableEntity,
		http.StatusInternalServerError,
	} {
		s, _ := newScriptedServer(t, "", status)
		_, e := newRequester(t, s.URL).Bytes(request.Get("/items"))
		assert.True(t, unexpected.Is(e))
	}
}

func TestNotFoundCarriesSitesReason(t *testing.T) {
	s, _ := newScriptedServer(
		t,
		`{"message":"no such item"}`,
		http.StatusNotFound,
	)
	_, e := newRequester(t, s.URL).Bytes(request.Get("/items/7"))
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "/items/7 status: 404: no such item", e.Error())
}

func TestClientRowOverridesTable(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusForbidden)
	_, e := newRequester(t, s.URL).
		WithStatus(http.StatusForbidden, errorsConstant.ForbiddenClass).
		Bytes(request.Get("/items"))
	assert.True(t, forbidden.Is(e))
}

func TestRequestRowOverridesClientRow(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusBadRequest)
	_, e := newRequester(t, s.URL).
		WithStatus(http.StatusBadRequest, errorsConstant.ForbiddenClass).
		Bytes(
			request.Get("/items").WithStatus(
				http.StatusBadRequest,
				errorsConstant.ValidationClass,
			),
		)
	assert.True(t, validation.Is(e))
}

func TestRefusalHookDecidesBeforeTable(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusConflict)
	_, e := newRequester(t, s.URL).
		WithRefusal(
			func(
				int,
				[]byte,
			) error {
				return validation.New("alfa")
			},
		).
		Bytes(request.Get("/items"))
	assert.True(t, validation.Is(e))
}

func TestSubjectAndReaderShapeMessage(t *testing.T) {
	s, _ := newScriptedServer(t, "<p>bravo</p>", http.StatusInternalServerError)
	_, e := newRequester(t, s.URL).
		WithSubject("alfa").
		WithReason(
			func(
				[]byte,
				string,
			) string {
				return "charlie"
			},
		).
		Bytes(request.Get("/items"))
	assert.True(t, unexpected.Is(e))
	assert.String(t, "alfa status: 500: charlie", e.Error())
}
