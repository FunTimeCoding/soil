package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestAnUnknownDecisionIsNotFound(t *testing.T) {
	s := webServer(t, service_tester.New(t).Service)
	status, _ := sendForm(t, s, decision_page.Path(999, true), nil)
	assert.Integer(t, http.StatusNotFound, status)
}

func TestAnOverlongReplyAlertsAtTheForm(t *testing.T) {
	v := service_tester.New(t).Service
	d, e := v.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, e)
	s := webServer(t, v)
	status, page := sendForm(
		t,
		s,
		fmt.Sprintf(
			"%s?%s=%d",
			constant.ReplyPath,
			constant.IdentifierParameter,
			d.Identifier,
		),
		url.Values{
			constant.ContentParameter: {
				strings.Repeat("x", constant.MaximumTurnLength+1),
			},
		},
	)
	assert.Integer(t, http.StatusOK, status)
	assert.StringContains(t, `role="alert"`, page)
	assert.StringContains(t, "the limit is 300", page)
}
