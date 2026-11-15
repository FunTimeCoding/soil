package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/prometheus_server"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/prometheus_tester"
	"net/http"
	"testing"
	"time"
)

func TestBadQueryIsCallersMistake(t *testing.T) {
	c := prometheus_tester.New(
		t,
		prometheus_server.New(
			t,
			http.StatusBadRequest,
			`{"status":"error","errorType":"bad_data","error":"parse error at char 4"}`,
		).URL,
	)
	_, e := c.Query("up{", time.Now())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "parse error at char 4", e.Error())
}

func TestQueryThatCannotExecuteIsCallersMistake(t *testing.T) {
	c := prometheus_tester.New(
		t,
		prometheus_server.New(
			t,
			http.StatusUnprocessableEntity,
			`{"status":"error","errorType":"execution","error":"many-to-many matching not allowed"}`,
		).URL,
	)
	_, e := c.Query("up * on() up", time.Now())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "many-to-many matching not allowed", e.Error())
}

func TestQueryTimeoutNeitherInvalidNorUnavailable(t *testing.T) {
	c := prometheus_tester.New(
		t,
		prometheus_server.New(
			t,
			http.StatusServiceUnavailable,
			`{"status":"error","errorType":"timeout","error":"query timed out in expression evaluation"}`,
		).URL,
	)
	_, e := c.Query(constant.Up, time.Now())
	assert.Error(t, e)
	assert.False(t, validation.Is(e))
	assert.False(t, unavailable.Is(e))
	assert.StringContains(t, "query timed out", e.Error())
}

func TestBadGatewayInFrontOfPrometheusIsUnavailable(t *testing.T) {
	c := prometheus_tester.New(
		t,
		prometheus_server.New(
			t,
			http.StatusBadGateway,
			"<html><body><h1>502 Bad Gateway</h1></body></html>",
		).URL,
	)
	_, e := c.Query(constant.Up, time.Now())
	assert.True(t, unavailable.Is(e))
}

func TestPathPrometheusLacksIsUnexpected(t *testing.T) {
	c := prometheus_tester.New(
		t,
		prometheus_server.New(t, http.StatusNotFound, "404 page not found").URL,
	)
	_, e := c.Query(constant.Up, time.Now())
	assert.True(t, unexpected.Is(e))
}

func TestRefusedPostFallsBackToGet(t *testing.T) {
	c := prometheus_tester.New(
		t,
		prometheus_server.New(t, http.StatusForbidden, "").URL,
	)
	result, e := c.Query(constant.Up, time.Now())
	assert.FatalOnError(t, e)
	assert.NotNil(t, result)
}
