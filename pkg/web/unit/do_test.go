package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/unit/failing_dialer"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestDoBytesReturnsBody(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer("alfa"))
	b, e := web.DoBytes(s.Client(), web.NewGet(join.Empty(s.URL, "/okay")))
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", string(b))
}

func TestDoClassifiesNotFound(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	_, e := web.Do(s.Client(), web.NewGet(join.Empty(s.URL, "/missing")))
	assert.True(t, not_found.Is(e))
}

func TestDoClassifiesOtherStatusAsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
	)
	_, e := web.Do(s.Client(), web.NewGet(join.Empty(s.URL, "/broken")))
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 500", e.Error())
}

func TestDoReadsLostHttp2ConnectionAsDropped(t *testing.T) {
	d := failing_dialer.New(errors.New(errorsConstant.ConnectionLost))
	_, e := web.Do(
		d.Client(),
		web.NewGet("https://foxtrot.example/wiki/rest/api/content/search"),
	)
	assert.True(t, dropped.Is(e))
	assert.String(t, "foxtrot.example: connection dropped", e.Error())
}

func TestDoLeavesQueryOutOfTransportFailures(t *testing.T) {
	s := upstream_server.New(t)
	locator := join.Empty(s.URL, "/okay?password=secret")
	s.Close()
	_, e := web.Do(s.Client(), web.NewGet(locator))
	assert.Error(t, e)
	assert.StringNotContains(t, "secret", e.Error())
}

func TestDoLeavesQueryOutOfMessages(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	_, e := web.Do(
		s.Client(),
		web.NewGet(join.Empty(s.URL, "/missing?api_key=secret")),
	)
	assert.StringNotContains(t, "secret", e.Error())
}
