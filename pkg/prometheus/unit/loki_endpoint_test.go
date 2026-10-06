package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestQueryRangeSendsTheWindow(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"status":"success","data":{"resultType":"streams","result":[]}}`,
	)
	start := time.Unix(100, 0)
	end := time.Unix(200, 0)
	_, e := newLokiClient(t, s.URL).QueryRange(
		`{namespace="alfa"}`,
		start,
		end,
		5,
	)
	assert.FatalOnError(t, e)
	r := last()
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
	s, _ := upstream_tester.New(t, http.StatusOK, `{"status":"error"}`)
	_, e := newLokiClient(t, s.URL).Labels(time.Unix(100, 0), time.Unix(200, 0))
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "loki status: error", e.Error())
}

func TestPlainTextRefusalKeepsTheReason(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusBadRequest,
		"parse error at line 1, col 1: syntax error: unexpected IDENTIFIER\n",
	)
	_, e := newLokiClient(t, s.URL).Query("charlie")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"loki status: 400: parse error at line 1, col 1: syntax error",
		e.Error(),
	)
}

func TestPushPostsTheStreams(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusNoContent, "")
	v := stream.New(map[string]string{"application": "delta"})
	v.Add(time.Unix(300, 0), "echo")
	assert.FatalOnError(t, newLokiClient(t, s.URL).Push(v))
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/loki/api/v1/push", r.URL.Path)
	assert.String(t, "application/json", r.Header.Get(constant.ContentType))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.StringContains(t, `"application":"delta"`, string(b))
	assert.StringContains(t, `"echo"`, string(b))
}

func TestMarkupRefusalFallsBackToTheStatus(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusBadGateway,
		"<html><body><h1>502 Bad Gateway</h1></body></html>",
	)
	e := newLokiClient(t, s.URL).Push(stream.New(map[string]string{}))
	assert.True(t, unexpected.Is(e))
	assert.StringNotContains(t, "<html>", e.Error())
	assert.StringContains(t, "status: 502", e.Error())
}
