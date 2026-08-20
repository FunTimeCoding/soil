package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http/httptest"
	"testing"
)

func TestParseEmpty(t *testing.T) {
	r := httptest.NewRequest(constant.Get, constant.LivePath, nil)
	s := subscription.Parse(r)
	assert.True(t, !s.Has("summary"))
}

func TestParseSingle(t *testing.T) {
	r := httptest.NewRequest(constant.Get, "/event?subscribe=summary", nil)
	s := subscription.Parse(r)
	assert.True(t, s.Has("summary"))
	assert.True(t, !s.Has("roster"))
}

func TestParseMultiple(t *testing.T) {
	r := httptest.NewRequest(
		constant.Get,
		"/event?subscribe=roster,activity",
		nil,
	)
	s := subscription.Parse(r)
	assert.True(t, s.Has("roster"))
	assert.True(t, s.Has("activity"))
	assert.True(t, !s.Has("summary"))
}

func TestQuerySingle(t *testing.T) {
	assert.String(t, "subscribe=summary", subscription.Query("summary"))
}

func TestQueryMultiple(t *testing.T) {
	assert.String(
		t,
		"subscribe=roster,activity",
		subscription.Query("roster", "activity"),
	)
}
