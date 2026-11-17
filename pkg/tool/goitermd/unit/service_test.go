package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/service_tester"
	"testing"
)

func TestReadHistoryDefaultsCount(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	v, e := o.Service().ReadHistory("alfa", 0)
	assert.FatalOnError(t, e)
	assert.Integer(t, 50, v.TotalScrollback)
}

func TestSendKeysSendsEveryKeyInOrder(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	assert.FatalOnError(
		t,
		o.Service().SendKeys("alfa", []string{"ctrl+c", "up", "return"}, 0),
	)
	assert.Strings(t, []string{"ctrl+c", "up", "return"}, o.Source.Keys("alfa"))
}

func TestSendKeysNamesFailingStep(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	e := o.Service().SendKeys("bravo", []string{"up"}, 0)
	assert.True(t, not_found.Is(e))
	assert.String(t, "step 1: session bravo not found", e.Error())
}

func TestSendKeysRequiresKeys(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	e := o.Service().SendKeys("alfa", nil, 0)
	assert.True(t, validation.Is(e))
}

func TestSendTextFollowsWithEnterWhenAsked(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	assert.FatalOnError(t, o.Service().SendText("alfa", "ls -la", true))
	assert.Strings(t, []string{"enter"}, o.Source.Keys("alfa"))
}

func TestSendTextRequiresText(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	e := o.Service().SendText("alfa", "", false)
	assert.True(t, validation.Is(e))
}

func TestSetTabTitleRequiresTitle(t *testing.T) {
	o := service_tester.New(t)
	o.AddSession("alfa")
	e := o.Service().SetTabTitle("tab-1", "")
	assert.True(t, validation.Is(e))
}
