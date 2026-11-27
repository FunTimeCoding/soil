package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/browser/firefox"
	"github.com/funtimecoding/soil/pkg/browser/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/browser/unit/extension_peer"
	"github.com/funtimecoding/soil/pkg/browser/unit/firefox_tester"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"testing"
	"time"
)

func TestClientAndItsMockShareFace(t *testing.T) {
	c := firefox.New()
	t.Cleanup(c.Close)
	m := mock_client.New(t)
	faces := []face.FirefoxClient{c, m}
	assert.Count(t, 2, faces)
}

func TestWithoutExtensionEveryCallIsUnavailable(t *testing.T) {
	c, _ := firefox_tester.New(t).Bridge()
	_, e := c.Tabs()
	assert.True(t, unavailable.Is(e))
	assert.False(t, classify.Reportable(e))
}

func TestCreatedTabComesBack(t *testing.T) {
	c, s := firefox_tester.New(t).Bridge()
	x := firefox_tester.New(t).Connected(c, s.URL)
	x.Answer(
		"create_tab",
		map[string]any{"tab_id": 7, "url": "https://example.com"},
	)
	v, e := c.CreateTab("https://example.com")
	assert.FatalOnError(t, e)
	assert.Integer(t, 7, v.Identifier)
}

func TestExtensionRefusalsCarryClasses(t *testing.T) {
	c, s := firefox_tester.New(t).Bridge()
	x := firefox_tester.New(t).Connected(c, s.URL)
	x.Refuse("click", "selector matched nothing: #missing")
	_, e := c.Click(1, "#missing")
	assert.True(t, not_found.Is(e))
	x.Refuse("wait_for", "selector not found within timeout: #late")
	_, f := c.WaitFor(1, "#late", 1)
	assert.True(t, not_found.Is(f))
	x.Refuse("evaluate", constant.FirefoxExpressionRequired)
	_, g := c.Evaluate(1, "")
	assert.True(t, validation.Is(g))
	x.Refuse("fill", constant.FirefoxSelectorRequired)
	_, k := c.Fill(1, "", "", false)
	assert.True(t, validation.Is(k))
	x.Refuse("download_file", constant.FirefoxLocatorRequired)
	_, l := c.DownloadFile("", "")
	assert.True(t, validation.Is(l))
	x.Refuse("read_tab", "Invalid tab ID: 9")
	_, h := c.ReadTab(9, false)
	assert.True(t, classify.Reportable(h))
}

func TestClosedBridgeRefusesExtension(t *testing.T) {
	c, s := firefox_tester.New(t).Bridge()
	firefox_tester.New(t).Connected(c, s.URL)
	c.Close()
	c.Close()
	_, e := c.Tabs()
	assert.True(t, unavailable.Is(e))
	extension_peer.Connect(t, s.URL)
	time.Sleep(50 * time.Millisecond)
	_, f := c.Tabs()
	assert.True(t, unavailable.Is(f))
}

func TestNewConnectionSurvivesOldOneEnding(t *testing.T) {
	c, s := firefox_tester.New(t).Bridge()
	first := firefox_tester.New(t).Connected(c, s.URL)
	second := extension_peer.Connect(t, s.URL)
	second.Answer("list_tabs", []any{})
	time.Sleep(50 * time.Millisecond)
	first.Disconnect()
	time.Sleep(50 * time.Millisecond)
	_, e := c.Tabs()
	assert.FatalOnError(t, e)
}
