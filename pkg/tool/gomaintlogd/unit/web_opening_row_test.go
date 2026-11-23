package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	"net/url"
	"testing"
)

func TestEntryRowOpensClosesSavesAndDeletesInPlace(t *testing.T) {
	s := base.New(t)
	s.PostBody(
		t,
		constant.AddEntryPath,
		url.Values{
			constant.Action:      {"restarted web server"},
			constant.User:        {"jdoe"},
			constant.System:      {"worker1"},
			constant.Service:     {"nginx"},
			constant.Description: {"nginx was unresponsive"},
		},
	)
	assert.StringContains(
		t,
		`<tr id="row-1-detail" hidden="hidden"></tr>`,
		s.Body(t, constant.EntriesPath),
	)
	opened := s.Body(t, "/detail?id=1")
	assert.StringContains(
		t,
		`<tr id="row-1-detail" class="opening-detail"><td colspan="6">`,
		opened,
	)
	assert.StringContains(
		t,
		`<template><tr id="row-1" class="opening-row" aria-expanded="true" hx-get="/close?id=1"`,
		opened,
	)
	closed := s.Body(t, "/close?id=1")
	assert.StringContains(
		t,
		`<tr id="row-1-detail" hidden="hidden"></tr>`,
		closed,
	)
	assert.StringContains(
		t,
		`aria-expanded="false" hx-get="/detail?id=1"`,
		closed,
	)
	saved := s.PostBody(
		t,
		"/edit?id=1",
		url.Values{
			constant.Action:      {"cleared and documented"},
			constant.User:        {"jdoe"},
			constant.System:      {"worker1"},
			constant.Service:     {"nginx"},
			constant.Description: {"documented"},
		},
	)
	assert.StringContains(t, `hx-swap-oob="true"><td class="time-cell">`, saved)
	assert.StringContains(
		t,
		`<td class="wrap-cell">cleared and documented</td>`,
		saved,
	)
	removed := s.PostBody(t, "/delete?id=1", nil)
	assert.StringContains(
		t,
		`<template><tr id="row-1" hx-swap-oob="delete"></tr></template>`,
		removed,
	)
	assert.StringContains(t, "Deleted entry", removed)
}
