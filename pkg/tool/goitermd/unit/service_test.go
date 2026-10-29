package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/service_tester"
	"testing"
)

func TestReadHistoryDefaultsTheCount(t *testing.T) {
	v, e := service_tester.New(t, sessionClient("alfa")).ReadHistory("alfa", 0)
	assert.FatalOnError(t, e)
	assert.Integer(t, 50, v.TotalScrollback)
}

func TestSendKeysSendsEveryKeyInOrder(t *testing.T) {
	c := sessionClient("alfa")
	assert.FatalOnError(
		t,
		service_tester.New(t, c).SendKeys(
			"alfa",
			[]string{"ctrl+c", "up", "return"},
			0,
		),
	)
	assert.Strings(t, []string{"ctrl+c", "up", "return"}, c.Keys("alfa"))
}

func TestSendKeysNamesTheFailingStep(t *testing.T) {
	e := service_tester.New(t, sessionClient("alfa")).SendKeys(
		"bravo",
		[]string{"up"},
		0,
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "step 1: session bravo not found", e.Error())
}

func TestSendKeysRequiresKeys(t *testing.T) {
	e := service_tester.New(t, sessionClient("alfa")).SendKeys("alfa", nil, 0)
	assert.True(t, validation.Is(e))
}

func TestSendTextFollowsWithEnterWhenAsked(t *testing.T) {
	c := sessionClient("alfa")
	assert.FatalOnError(
		t,
		service_tester.New(t, c).SendText("alfa", "ls -la", true),
	)
	assert.Strings(t, []string{"enter"}, c.Keys("alfa"))
}

func TestSendTextRequiresText(t *testing.T) {
	e := service_tester.New(t, sessionClient("alfa")).SendText(
		"alfa",
		"",
		false,
	)
	assert.True(t, validation.Is(e))
}

func TestSetTabTitleRequiresATitle(t *testing.T) {
	e := service_tester.New(t, sessionClient("alfa")).SetTabTitle("tab-1", "")
	assert.True(t, validation.Is(e))
}
