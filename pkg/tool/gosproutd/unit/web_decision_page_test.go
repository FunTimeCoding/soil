package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"strings"
	"testing"
)

func TestTheGridPreservesEveryCard(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		[]string{"narrow", "wide"},
		nil,
	)
	assert.FatalOnError(t, e)
	page := renderNode(
		t,
		session_page.Grid(s.Service.Decisions(stringsConstant.LowerAlfa)),
	)
	assert.StringContains(
		t,
		fmt.Sprintf(`id="decision-%d" hx-preserve="true"`, d.Identifier),
		page,
	)
	assert.True(
		t,
		!strings.Contains(renderNode(t, decision_page.Card(d)), "hx-preserve"),
	)
}

func TestAnOpenedCardKeepsItsFormsOutsideTheLiveRegions(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		[]string{"narrow", "wide"},
		nil,
	)
	assert.FatalOnError(t, e)
	page := renderNode(t, decision_page.Hero(d, "", ""))
	markup_tester.AssertCount(t, 2, page, `input[type="text"]`)
	markup_tester.AssertCount(t, 0, page, `[sse-swap] input[type="text"]`)
	markup_tester.AssertCount(t, 4, page, "[sse-swap]")
}
