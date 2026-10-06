package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestGenerateReturnsTheReport(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, "bravo report")
	result, e := newClient(t, s.URL).Generate([]string{"charlie.json"}, nil)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo report", result)
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "Bearer alfa-token", r.Header.Get(constant.Authorization))
}

func TestFailedGenerateIsAnError(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusInternalServerError,
		"delta parse failed",
	)
	_, e := newClient(t, s.URL).Generate([]string{"charlie.json"}, nil)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"raid parser generate status: 500: delta parse failed",
		e.Error(),
	)
}
