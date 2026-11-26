package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/unit/service_tester"
	"os"
	"testing"
)

func TestNavigateTargetsFirstTab(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Client.AddContext("bravo", "about:blank")
	assert.FatalOnError(t, o.Service().Navigate("", "https://alfa.example/"))
	tabs, e := o.Client.Tree()
	assert.FatalOnError(t, e)
	assert.String(t, "https://alfa.example/", tabs[0].Locator)
	assert.String(t, "about:blank", tabs[1].Locator)
}

func TestNavigateTargetsNamedTab(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Client.AddContext("bravo", "about:blank")
	assert.FatalOnError(
		t,
		o.Service().Navigate("bravo", "https://bravo.example/"),
	)
	tabs, e := o.Client.Tree()
	assert.FatalOnError(t, e)
	assert.String(t, "about:blank", tabs[0].Locator)
	assert.String(t, "https://bravo.example/", tabs[1].Locator)
}

func TestNavigateUnknownTabNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	e := o.Service().Navigate("charlie", "https://charlie.example/")
	assert.True(t, not_found.Is(e))
	assert.String(t, "tab not found: charlie", e.Error())
}

func TestNavigateRequiresLocator(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	assert.True(t, validation.Is(o.Service().Navigate("", "")))
}

func TestNavigateWithoutTabsValidation(t *testing.T) {
	e := service_tester.New(t).Service().Navigate("", "https://alfa.example/")
	assert.True(t, validation.Is(e))
}

func TestEvaluateReturnsResult(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Client.SetResult("1+2", `{"type":"number","value":3}`)
	result, e := o.Service().Evaluate("", "1+2")
	assert.FatalOnError(t, e)
	assert.String(t, `{"type":"number","value":3}`, string(result))
}

func TestEvaluateRequiresExpression(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	_, e := o.Service().Evaluate("", "")
	assert.True(t, validation.Is(e))
}

func TestScreenshotSavesImage(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	path, e := o.Service().Screenshot("")
	assert.FatalOnError(t, e)
	assert.Suffix(t, "ladybird_alfa.png", path)
	b, f := os.ReadFile(path)
	assert.FatalOnError(t, f)
	assert.String(t, "png", string(b))
}
