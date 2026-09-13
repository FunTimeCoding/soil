package web_interface

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration/web_interface_tester"
	"strings"
	"testing"
)

func TestDashboardIsLive(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	body := o.Get(constant.DashboardPath)
	assert.StringContains(t, "sse-connect", body)
	assert.StringContains(t, "summary_strip", body)
	assert.StringContains(t, "recent", body)
}

func TestDashboardNoLongerPolls(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	assert.True(
		t,
		!strings.Contains(o.Get(constant.DashboardPath), "every 60s"),
	)
}
