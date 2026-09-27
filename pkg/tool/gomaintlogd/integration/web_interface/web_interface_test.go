package web_interface

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration/web_interface_tester"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestFilterWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	o.PostForm(
		constant.AddEntryPath,
		url.Values{
			"action":      {"restart"},
			"user":        {"alice"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"test"},
		},
	)
	o.PostForm(
		constant.AddEntryPath,
		url.Values{
			"action":      {"backup"},
			"user":        {"bob"},
			"system":      {"worker2"},
			"service":     {"nginx"},
			"description": {"test"},
		},
	)
	b := o.Get("/entries?system=worker1")
	assert.StringContains(t, "restart", b)
	assert.StringContains(t, "worker1", b)
	assert.StringContains(
		t,
		"backup",
		o.Get(fmt.Sprintf("/entries?user=%s", "bob")),
	)
	assert.StringContains(
		t,
		"restart",
		o.Get(fmt.Sprintf("/entries?user=%s", "alice")),
	)
	assert.StringContains(
		t,
		"No entries found",
		o.Get("/entries?system=nonexistent"),
	)
}

func TestDashboardIsLive(t *testing.T) {
	o := web_interface_tester.New(t)
	body := o.Get(constant.DashboardPath)
	assert.StringContains(t, "sse-connect", body)
	assert.StringContains(t, "summary_strip", body)
	assert.StringContains(t, "recent", body)
}

func TestDashboardNoLongerPolls(t *testing.T) {
	o := web_interface_tester.New(t)
	assert.True(
		t,
		!strings.Contains(o.Get(constant.DashboardPath), "every 60s"),
	)
}

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertStatus(http.StatusOK, constant.DashboardPath)
	o.AssertStatus(http.StatusOK, constant.EntriesPath)
	o.AssertStatus(http.StatusOK, constant.AddEntryPath)
	assert.StringContains(t, "No entries found", o.Get(constant.DashboardPath))
	addBody := o.PostForm(
		constant.AddEntryPath,
		url.Values{
			"action":      {"restarted web server"},
			"user":        {"jdoe"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"nginx was unresponsive, restarted"},
		},
	)
	assert.StringContains(t, "Entry added", addBody)
	assert.StringContains(t, "restarted web server", addBody)
	assert.StringContains(
		t,
		"restarted web server",
		o.Get(constant.DashboardPath),
	)
	assert.StringContains(t, "jdoe", o.Get(constant.DashboardPath))
	assert.StringContains(t, "worker1", o.Get(constant.EntriesPath))
	detail := o.Get("/detail?id=1")
	assert.StringContains(t, "nginx was unresponsive, restarted", detail)
	assert.StringContains(t, "Edit", detail)
	assert.StringContains(t, "Delete", detail)
	assert.StringContains(t, "Permalink", detail)
	assert.StringContains(t, `href="/entry/1"`, detail)
	o.AssertStatus(http.StatusOK, "/entry/1")
	page := o.Get("/entry/1")
	assert.StringContains(t, "nginx was unresponsive, restarted", page)
	assert.StringContains(t, "worker1", page)
	o.AssertStatus(http.StatusNotFound, "/entry/9999")
	assert.StringContains(
		t,
		"This entry no longer exists",
		o.Get("/entry/9999"),
	)
	o.AssertStatus(http.StatusNotFound, "/entry/nonsense")
	edit := o.Get("/edit?id=1")
	assert.StringContains(t, "restarted web server", edit)
	assert.StringContains(t, "Save", edit)
	editBody := o.PostForm(
		"/edit?id=1",
		url.Values{
			"action":      {"cleared and documented"},
			"user":        {"jdoe"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"nginx was unresponsive, restarted and documented"},
			"timestamp":   {"2026-03-19T10:00"},
		},
	)
	assert.StringContains(t, "cleared and documented", editBody)
	o.PostForm("/delete?id=1", nil)
	assert.StringContains(t, "No entries found", o.Get(constant.DashboardPath))
}
