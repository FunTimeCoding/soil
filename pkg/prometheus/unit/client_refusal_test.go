package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"net/http"
	"testing"
	"time"
)

func TestABadQueryIsTheCallersMistake(t *testing.T) {
	c := newRefusingClient(
		t,
		http.StatusBadRequest,
		`{"status":"error","errorType":"bad_data","error":"parse error at char 4"}`,
	)
	_, e := c.Query("up{", time.Now())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "parse error at char 4", e.Error())
}

func TestAQueryThatCannotExecuteIsTheCallersMistake(t *testing.T) {
	c := newRefusingClient(
		t,
		http.StatusUnprocessableEntity,
		`{"status":"error","errorType":"execution","error":"many-to-many matching not allowed"}`,
	)
	_, e := c.Query("up * on() up", time.Now())
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "many-to-many matching not allowed", e.Error())
}

func TestAQueryTimeoutIsNotAVerdictOnThePeer(t *testing.T) {
	c := newRefusingClient(
		t,
		http.StatusServiceUnavailable,
		`{"status":"error","errorType":"timeout","error":"query timed out in expression evaluation"}`,
	)
	_, e := c.Query(constant.Up, time.Now())
	assert.Error(t, e)
	assert.False(t, validation.Is(e))
	assert.False(t, unavailable.Is(e))
	assert.StringContains(t, "query timed out", e.Error())
}

func TestARefusedPostFallsBackToGet(t *testing.T) {
	c := newRefusingClient(t, http.StatusForbidden, "")
	result, e := c.Query(constant.Up, time.Now())
	assert.FatalOnError(t, e)
	assert.NotNil(t, result)
}
