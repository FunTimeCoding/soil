package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/query"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"github.com/funtimecoding/soil/pkg/web/types/renewing_authorizer"
	"net/http"
	"net/url"
	"testing"
)

func TestNotationSendsAgentHeaderAndAuthorization(t *testing.T) {
	s, seen := newScriptedServer(t, `{"name":"alfa"}`, http.StatusOK)
	var out named
	e := newRequester(t, s.URL).
		WithUserAgent("bravo/1.0").
		WithHeader(constant.Accept, "text/css").
		WithAuthorizer(basic.New("charlie", "delta")).
		Notation(request.Get("/items").WithParameter("page", "2"), &out)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", out.Name)
	r := seen()[0]
	assert.String(t, "/items", r.URL.Path)
	assert.String(t, "2", r.URL.Query().Get("page"))
	assert.String(t, "bravo/1.0", r.Header.Get(constant.UserAgent))
	assert.String(t, "text/css", r.Header.Get(constant.Accept))
	user, password, _ := r.BasicAuth()
	assert.String(t, "charlie", user)
	assert.String(t, "delta", password)
}

func TestMissingIsNotFound(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusNotFound)
	_, e := newRequester(t, s.URL).Bytes(request.Get("/missing"))
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesTheSitesReason(t *testing.T) {
	s, _ := newScriptedServer(
		t,
		`{"message":"quota spent"}`,
		http.StatusBadRequest,
	)
	_, e := newRequester(t, s.URL).Bytes(request.Get("/items"))
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 400: quota spent", e.Error())
}

func TestRefusalCanBeReadBySite(t *testing.T) {
	s, _ := newScriptedServer(t, "spent", http.StatusBadRequest)
	_, e := newRequester(t, s.URL).
		WithRefusal(
			func(
				status int,
				body []byte,
			) error {
				return validation.New("site says %s", body)
			},
		).
		Bytes(request.Get("/items"))
	assert.True(t, validation.Is(e))
	assert.String(t, "site says spent", e.Error())
}

func TestTransportFailureHidesTheKey(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusOK)
	r := newRequester(t, s.URL).WithAuthorizer(
		query.New(url.Values{"api_key": {"secret"}}),
	)
	s.Close()
	_, e := r.Bytes(request.Get("/items"))
	assert.True(t, unreachable.Is(e))
	assert.StringNotContains(t, "secret", e.Error())
}

func TestGetRetriesTransientStatus(t *testing.T) {
	s, seen := newScriptedServer(
		t,
		"alfa",
		http.StatusServiceUnavailable,
		http.StatusOK,
	)
	b, e := newRequester(t, s.URL).Bytes(request.Get("/items"))
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", string(b))
	assert.Count(t, 2, seen())
}

func TestPostIsSentOnce(t *testing.T) {
	s, seen := newScriptedServer(t, "", http.StatusServiceUnavailable)
	_, e := newRequester(t, s.URL).Bytes(
		request.PostForm("/items", url.Values{"name": {"alfa"}}),
	)
	assert.True(t, unexpected.Is(e))
	assert.Count(t, 1, seen())
}

func TestUnauthorizedRenewsOnce(t *testing.T) {
	a := renewing_authorizer.New("stale")
	var out named
	e := newRequester(t, newRenewingServer(t, "Bearer fresh").URL).
		WithAuthorizer(a).
		Notation(request.Get("/items"), &out)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", out.Name)
	assert.Integer(t, 1, a.Renewed)
}

func TestSecondUnauthorizedIsReturned(t *testing.T) {
	a := renewing_authorizer.New("stale")
	_, e := newRequester(t, newRenewingServer(t, "Bearer never").URL).
		WithAuthorizer(a).
		Bytes(request.Get("/items"))
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 401", e.Error())
	assert.Integer(t, 1, a.Renewed)
}

func TestAbsoluteLocatorLeavesTheBase(t *testing.T) {
	s, seen := newScriptedServer(t, "alfa", http.StatusOK)
	b, e := requester.New(locator.New("unused.example")).Bytes(
		request.Absolute(join.Empty(s.URL, "/files/a.png")),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", string(b))
	assert.String(t, "/files/a.png", seen()[0].URL.Path)
}

func TestAbsoluteLocatorKeepsItsParameters(t *testing.T) {
	s, seen := newScriptedServer(t, "alfa", http.StatusOK)
	_, e := requester.New(locator.New("unused.example")).Bytes(
		request.Absolute(join.Empty(s.URL, "/files/a.png?size=full")).
			WithParameter("sid", "bravo"),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "full", seen()[0].URL.Query().Get("size"))
	assert.String(t, "bravo", seen()[0].URL.Query().Get("sid"))
}
