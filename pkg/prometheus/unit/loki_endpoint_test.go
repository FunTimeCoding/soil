package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/loki_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestQueryRangeSendsWindow(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"status":"success","data":{"resultType":"streams","result":[]}}`,
		),
	)
	start := time.Unix(100, 0)
	end := time.Unix(200, 0)
	_, e := loki_tester.New(t, s.URL).QueryRange(
		`{namespace="alfa"}`,
		start,
		end,
		5,
	)
	assert.FatalOnError(t, e)
	r := s.Last()
	assert.String(t, "/loki/api/v1/query_range", r.URL.Path)
	assert.String(t, `{namespace="alfa"}`, r.URL.Query().Get("query"))
	assert.String(t, "100000000000", r.URL.Query().Get("start"))
	assert.String(t, "200000000000", r.URL.Query().Get("end"))
	assert.String(t, "5", r.URL.Query().Get("limit"))
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo-password", password)
}

func TestFailedStatusIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"status":"error"}`),
	)
	_, e := loki_tester.New(t, s.URL).Labels(
		time.Unix(100, 0),
		time.Unix(200, 0),
	)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "loki status: error", e.Error())
}

func TestPlainTextRefusalKeepsReason(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			"parse error at line 1, col 1: syntax error: unexpected IDENTIFIER\n",
		),
	)
	_, e := loki_tester.New(t, s.URL).Query("charlie")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"loki status: 400: parse error at line 1, col 1: syntax error",
		e.Error(),
	)
}

func TestPushPostsStreams(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNoContent),
	)
	v := stream.New(map[string]string{"application": "delta"})
	v.Add(time.Unix(300, 0), "echo")
	assert.FatalOnError(t, loki_tester.New(t, s.URL).Push(v))
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/loki/api/v1/push", r.URL.Path)
	assert.String(t, "application/json", r.Header.Get(webConstant.ContentType))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.StringContains(t, `"application":"delta"`, string(b))
	assert.StringContains(t, `"echo"`, string(b))
}

func TestMarkupRefusalFallsBackToStatus(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadGateway),
		upstream_server.WithAnswer(
			"<html><body><h1>502 Bad Gateway</h1></body></html>",
		),
	)
	e := loki_tester.New(t, s.URL).Push(stream.New(map[string]string{}))
	assert.True(t, unavailable.Is(e))
	assert.StringNotContains(t, "<html>", e.Error())
	assert.StringContains(t, "status: 502", e.Error())
}
