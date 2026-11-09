package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/service_tester"
	"testing"
)

func TestReadHistoryDefaultsCount(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	v, e := service_tester.New(t, c).ReadHistory("alfa", 0)
	assert.FatalOnError(t, e)
	assert.Integer(t, 50, v.TotalScrollback)
}

func TestSendKeysSendsEveryKeyInOrder(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
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

func TestSendKeysNamesFailingStep(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	e := service_tester.New(t, c).SendKeys("bravo", []string{"up"}, 0)
	assert.True(t, not_found.Is(e))
	assert.String(t, "step 1: session bravo not found", e.Error())
}

func TestSendKeysRequiresKeys(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	e := service_tester.New(t, c).SendKeys("alfa", nil, 0)
	assert.True(t, validation.Is(e))
}

func TestSendTextFollowsWithEnterWhenAsked(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	assert.FatalOnError(
		t,
		service_tester.New(t, c).SendText("alfa", "ls -la", true),
	)
	assert.Strings(t, []string{"enter"}, c.Keys("alfa"))
}

func TestSendTextRequiresText(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	e := service_tester.New(t, c).SendText("alfa", "", false)
	assert.True(t, validation.Is(e))
}

func TestSetTabTitleRequiresTitle(t *testing.T) {
	c := sessionClient("alfa")
	defer c.Close()
	e := service_tester.New(t, c).SetTabTitle("tab-1", "")
	assert.True(t, validation.Is(e))
}
