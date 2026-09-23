package web_interface

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/integration/web_interface_tester"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertStatus(http.StatusOK, constant.DashboardPath)
	o.AssertStatus(http.StatusOK, constant.RecentPath)
	o.AssertStatus(http.StatusOK, "/alerts?name=HighMemory")
}

func TestDashboardIsLive(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertContains("sse-connect", constant.DashboardPath)
	o.AssertContains("summary_strip", constant.DashboardPath)
	o.AssertContains("top", constant.DashboardPath)
}

func TestDashboardNoLongerPolls(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertMissing("every 60s", constant.DashboardPath)
}
