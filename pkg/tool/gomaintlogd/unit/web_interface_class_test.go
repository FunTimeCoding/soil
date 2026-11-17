package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	"net/http"
	"net/url"
	"testing"
)

func TestWebInterfaceAnswersMalformedEntryQuietly(t *testing.T) {
	s := base.New(t)
	assert.Integer(t, http.StatusBadRequest, s.Status(t, "/detail?id=nonsense"))
	assert.Integer(t, http.StatusBadRequest, s.Status(t, "/edit?id=0"))
	assert.Integer(
		t,
		http.StatusBadRequest,
		s.PostStatus(t, "/delete?id=nonsense", nil),
	)
	assert.Integer(t, 0, s.Captured())
}

func TestWebInterfaceAnswersUnknownEntryQuietly(t *testing.T) {
	s := base.New(t)
	assert.Integer(t, http.StatusNotFound, s.Status(t, "/detail?id=9999"))
	assert.Integer(t, http.StatusNotFound, s.Status(t, "/edit?id=9999"))
	assert.Integer(
		t,
		http.StatusNotFound,
		s.PostStatus(
			t,
			"/edit?id=9999",
			url.Values{constant.Action: {"x"}, constant.User: {"y"}},
		),
	)
	assert.Integer(
		t,
		http.StatusNotFound,
		s.PostStatus(t, "/delete?id=9999", nil),
	)
	assert.Integer(t, 0, s.Captured())
}

func TestWebInterfaceRefusesEntryWithoutUser(t *testing.T) {
	s := base.New(t)
	assert.StringContains(
		t,
		"user is required",
		s.PostBody(
			t,
			constant.AddEntryPath,
			url.Values{constant.Action: {"restart"}},
		),
	)
	assert.Integer(t, 0, s.Captured())
}
