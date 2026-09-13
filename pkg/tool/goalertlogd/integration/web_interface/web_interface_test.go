package web_interface

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/integration/web_interface_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertStatus(constant.DashboardPath, http.StatusOK)
	o.AssertStatus(constant.RecentPath, http.StatusOK)
	o.AssertStatus("/alerts?name=HighMemory", http.StatusOK)
}

func TestDashboardIsLive(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertContains(constant.DashboardPath, "sse-connect")
	o.AssertContains(constant.DashboardPath, webConstant.LayoutSummaryStrip)
	o.AssertContains(constant.DashboardPath, constant.EventTop)
}

func TestDashboardNoLongerPolls(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertMissing(constant.DashboardPath, "every 60s")
}
