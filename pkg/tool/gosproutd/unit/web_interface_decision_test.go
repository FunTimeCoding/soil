package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/web_interface_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestUnknownDecisionIsNotFound(t *testing.T) {
	s, _ := base.New(t)
	status, _ := web_interface_tester.Send(
		t,
		s,
		decision_page.Path(999, true),
		nil,
	)
	assert.Integer(t, http.StatusNotFound, status)
}

func TestOverlongReplyAlertsAtForm(t *testing.T) {
	s, seed := base.New(t)
	d := seed.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	status, page := web_interface_tester.Send(
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
