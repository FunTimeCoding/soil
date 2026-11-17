package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	"net/http"
	"net/url"
	"testing"
)

func TestFilterWebInterface(t *testing.T) {
	s := base.New(t)
	s.PostBody(
		t,
		constant.AddEntryPath,
		url.Values{
			"action":      {"restart"},
			"user":        {"alice"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"test"},
		},
	)
	s.PostBody(
		t,
		constant.AddEntryPath,
		url.Values{
			"action":      {"backup"},
			"user":        {"bob"},
			"system":      {"worker2"},
			"service":     {"nginx"},
			"description": {"test"},
		},
	)
	b := s.Body(t, "/entries?system=worker1")
	assert.StringContains(t, "restart", b)
	assert.StringContains(t, "worker1", b)
	assert.StringContains(
		t,
		"backup",
		s.Body(t, fmt.Sprintf("/entries?user=%s", "bob")),
	)
	assert.StringContains(
		t,
		"restart",
		s.Body(t, fmt.Sprintf("/entries?user=%s", "alice")),
	)
	assert.StringContains(
		t,
		"No entries found",
		s.Body(t, "/entries?system=nonexistent"),
	)
}

func TestDashboardIsLive(t *testing.T) {
	body := base.New(t).Body(t, constant.DashboardPath)
	assert.StringContains(t, "sse-connect", body)
	assert.StringContains(t, "summary_strip", body)
	assert.StringContains(t, "recent", body)
}

func TestDashboardNoLongerPolls(t *testing.T) {
	assert.StringNotContains(
		t,
		"every 60s",
		base.New(t).Body(t, constant.DashboardPath),
	)
}

func TestWebInterface(t *testing.T) {
	s := base.New(t)
	assert.Integer(t, http.StatusOK, s.Status(t, constant.DashboardPath))
	assert.Integer(t, http.StatusOK, s.Status(t, constant.EntriesPath))
	assert.Integer(t, http.StatusOK, s.Status(t, constant.AddEntryPath))
	assert.StringContains(
		t,
		"No entries found",
		s.Body(t, constant.DashboardPath),
	)
	addBody := s.PostBody(
		t,
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
		s.Body(t, constant.DashboardPath),
	)
	assert.StringContains(t, "jdoe", s.Body(t, constant.DashboardPath))
	assert.StringContains(t, "worker1", s.Body(t, constant.EntriesPath))
	detail := s.Body(t, "/detail?id=1")
	assert.StringContains(t, "nginx was unresponsive, restarted", detail)
	assert.StringContains(t, "Edit", detail)
	assert.StringContains(t, "Delete", detail)
	assert.StringContains(t, "Permalink", detail)
	assert.StringContains(t, `href="/entry/1"`, detail)
	assert.Integer(t, http.StatusOK, s.Status(t, "/entry/1"))
	page := s.Body(t, "/entry/1")
	assert.StringContains(t, "nginx was unresponsive, restarted", page)
	assert.StringContains(t, "worker1", page)
	assert.Integer(t, http.StatusNotFound, s.Status(t, "/entry/9999"))
	assert.StringContains(
		t,
		"This entry no longer exists",
		s.Body(t, "/entry/9999"),
	)
	assert.Integer(t, http.StatusNotFound, s.Status(t, "/entry/nonsense"))
	edit := s.Body(t, "/edit?id=1")
	assert.StringContains(t, "restarted web server", edit)
	assert.StringContains(t, "Save", edit)
	editBody := s.PostBody(
		t,
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
	s.PostBody(t, "/delete?id=1", nil)
	assert.StringContains(
		t,
		"No entries found",
		s.Body(t, constant.DashboardPath),
	)
}
