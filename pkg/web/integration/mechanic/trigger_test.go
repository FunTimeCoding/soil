//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestTriggerLoad(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	b.WaitCondition(
		"document.querySelector('#load-region').textContent.indexOf('term=load') >= 0",
	)
	assert.StringContains(t, "term=load", b.Text("#load-region"))
}

func TestTriggerLoadDelay(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	assert.String(t, "", b.Text("#delayed"))
	b.WaitCondition(
		"document.querySelector('#delayed').textContent.indexOf('term=delayed') >= 0",
	)
}

func TestTriggerChange(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	b.WaitVisible("#scope-select")
	b.Evaluate(
		"(function(){var s=document.querySelector('#scope-select');s.value='second';s.dispatchEvent(new Event('change'));})()",
		nil,
	)
	b.WaitCondition(
		"document.querySelector('#selection').textContent.indexOf('scope=second') >= 0",
	)
}

func TestTriggerKeyupIncludesAndVals(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.TriggerPath))
	b.WaitVisible("#term-input")
	b.Fill("#term-input", "abc")
	b.WaitCondition(
		"document.querySelector('#search').textContent.indexOf('term=abc') >= 0",
	)
	text := b.Text("#search")
	assert.StringContains(t, "term=abc", text)
	assert.StringContains(t, "scope=first", text)
	assert.StringContains(t, "origin=typed", text)
}
