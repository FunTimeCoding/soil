package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/web_interface_tester"
	"net/http"
	"net/url"
	"testing"
)

func TestWebInterfaceAnswersAMalformedEntryQuietly(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertStatus(http.StatusBadRequest, "/detail?id=nonsense")
	o.AssertStatus(http.StatusBadRequest, "/edit?id=0")
	o.AssertPostStatus(http.StatusBadRequest, "/delete?id=nonsense", nil)
	assert.Integer(t, 0, o.Captured())
}

func TestWebInterfaceAnswersAnUnknownEntryQuietly(t *testing.T) {
	o := web_interface_tester.New(t)
	o.AssertStatus(http.StatusNotFound, "/detail?id=9999")
	o.AssertStatus(http.StatusNotFound, "/edit?id=9999")
	o.AssertPostStatus(
		http.StatusNotFound,
		"/edit?id=9999",
		url.Values{constant.Action: {"x"}, constant.User: {"y"}},
	)
	o.AssertPostStatus(http.StatusNotFound, "/delete?id=9999", nil)
	assert.Integer(t, 0, o.Captured())
}

func TestWebInterfaceRefusesAnEntryWithoutUser(t *testing.T) {
	o := web_interface_tester.New(t)
	assert.StringContains(
		t,
		"user is required",
		o.PostForm(
			constant.AddEntryPath,
			url.Values{constant.Action: {"restart"}},
		),
	)
	assert.Integer(t, 0, o.Captured())
}
