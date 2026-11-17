package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	s := base.New(t)
	assert.Integer(t, http.StatusOK, s.Status(t, constant.DashboardPath))
	assert.Integer(t, http.StatusOK, s.Status(t, constant.RecentPath))
	assert.Integer(t, http.StatusOK, s.Status(t, "/alerts?name=HighMemory"))
}

func TestDashboardIsLive(t *testing.T) {
	body := base.New(t).Body(t, constant.DashboardPath)
	assert.StringContains(t, "sse-connect", body)
	assert.StringContains(t, "summary_strip", body)
	assert.StringContains(t, "top", body)
}

func TestDashboardNoLongerPolls(t *testing.T) {
	assert.StringNotContains(
		t,
		"every 60s",
		base.New(t).Body(t, constant.DashboardPath),
	)
}
