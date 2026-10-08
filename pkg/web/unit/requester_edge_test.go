package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
	"testing"
)

func TestCredentialsStayOnTheirHost(t *testing.T) {
	own, ownSeen := newScriptedServer(t, "alfa", http.StatusOK)
	other, otherSeen := newScriptedServer(t, "bravo", http.StatusOK)
	r := newRequester(t, own.URL).WithAuthorizer(basic.New("charlie", "delta"))
	_, e := r.Bytes(request.Get("/items"))
	assert.FatalOnError(t, e)
	_, e = r.Bytes(request.Absolute(join.Empty(other.URL, "/files/a.png")))
	assert.FatalOnError(t, e)
	assert.True(t, ownSeen()[0].Header.Get(constant.Authorization) != "")
	assert.String(t, "", otherSeen()[0].Header.Get(constant.Authorization))
}

func TestUnknownHostIsNotRetried(t *testing.T) {
	client, attempts := newFailingClient(
		&net.DNSError{Err: "no such host", Name: "echo.example"},
	)
	_, e := requester.New(locator.New("echo.example")).
		WithClient(client).
		WithBackoff(0).
		Bytes(request.Get("/items"))
	assert.True(t, unreachable.Is(e))
	assert.Integer(t, 1, int(attempts.Load()))
}

func TestRefusedConnectionIsRetried(t *testing.T) {
	client, attempts := newFailingClient(
		os.NewSyscallError("connect", syscall.ECONNREFUSED),
	)
	_, e := requester.New(locator.New("echo.example")).
		WithClient(client).
		WithBackoff(0).
		Bytes(request.Get("/items"))
	assert.True(t, unreachable.Is(e))
	assert.Integer(t, 3, int(attempts.Load()))
}

func TestPatchSendsNotation(t *testing.T) {
	s, seen := newScriptedServer(t, "{}", http.StatusOK)
	_, e := newRequester(t, s.URL).Bytes(
		request.New(http.MethodPatch, "/items/1").WithNotation(
			map[string]string{"name": "alfa"},
		),
	)
	assert.FatalOnError(t, e)
	r := seen()[0]
	assert.String(t, http.MethodPatch, r.Method)
	assert.String(t, "application/json", r.Header.Get(constant.ContentType))
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)
	assert.String(t, `{"name":"alfa"}`, string(b))
}

func TestUntrustedCertificateIsUnreachable(t *testing.T) {
	s := httptest.NewTLSServer(http.NotFoundHandler())
	t.Cleanup(s.Close)
	_, e := requester.New(locator.New("unused.example")).
		Bytes(request.Absolute(s.URL))
	assert.True(t, unreachable.Is(e))
	assert.StringContains(t, "untrusted certificate", e.Error())
	_, e = requester.New(locator.New("unused.example")).
		WithClient(web.InsecureStallClient()).
		Bytes(request.Absolute(s.URL))
	assert.True(t, !unreachable.Is(e))
}

func TestRefusalReadsCommonShapes(t *testing.T) {
	for body, expect := range map[string]string{
		`{"detail":"alfa"}`:                  "alfa",
		`{"errorMessages":["bravo"]}`:        "bravo",
		`{"error":{"message":"charlie"}}`:    "charlie",
		`{"errors":[{"title":"delta"}]}`:     "delta",
		`{"errors":[{"detail":"echo"}]}`:     "echo",
		`{"errorMessage":"foxtrot"}`:         "foxtrot",
		`{"title":"golf"}`:                   "golf",
		`{"title":"hotel","detail":"india"}`: "india",
	} {
		s, _ := newScriptedServer(t, body, http.StatusBadRequest)
		_, e := newRequester(t, s.URL).Bytes(request.Get("/items"))
		assert.StringContains(t, join.Empty("status: 400: ", expect), e.Error())
	}
}

func TestAnswerThatIsNotNotationIsUnexpected(t *testing.T) {
	s, _ := newScriptedServer(t, "<html>", http.StatusOK)
	var out Named
	e := newRequester(t, s.URL).Notation(request.Get("/items"), &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "/items: answer is not JSON", e.Error())
}

func TestNoContentSucceeds(t *testing.T) {
	s, _ := newScriptedServer(t, "", http.StatusNoContent)
	b, e := newRequester(t, s.URL).Bytes(
		request.New(http.MethodDelete, "/items/1"),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, b)
}
