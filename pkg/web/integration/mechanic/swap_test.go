//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestSwapInner(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#counter-post")
	assert.String(t, "0", b.Text("#counter"))
	b.Click("#counter-post")
	b.WaitCondition("document.querySelector('#counter').textContent === '1'")
	assert.String(t, "1", b.Text("#counter"))
}

func TestSwapGetReadsWithoutIncrement(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#counter-get")
	b.Click("#counter-get")
	b.WaitCondition("document.querySelector('#counter') !== null")
	assert.String(t, "0", b.Text("#counter"))
}

func TestSwapOuterReplacesElement(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#row-post")
	assert.String(t, "original", b.Text("#row"))
	b.Evaluate("window.confirm = function() { return true; }", nil)
	b.Click("#row-post")
	b.WaitCondition("document.querySelector('#row').textContent === 'replaced'")
	assert.Integer(t, 1, b.CountElements("#row"))
}

func TestSwapConfirmCancelBlocksRequest(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.RootPath))
	b.WaitVisible("#row-post")
	b.Evaluate("window.confirm = function() { return false; }", nil)
	b.Click("#row-post")
	b.WaitCondition("true")
	assert.String(t, "original", b.Text("#row"))
}
