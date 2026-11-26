//go:build browser

package browser

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/integration/browser/browser_tester"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/fill"
	"testing"
)

func TestServiceDrivesPage(t *testing.T) {
	s := browser_tester.New(t).Service()
	assert.FatalOnError(
		t,
		s.Navigate("", browser_tester.Page(constant.FixturePage)),
	)
	v, e := s.Snapshot("")
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"[document] Probe\n  uid=e1 [button] Press\n  uid=e2 [label] Name\n    uid=e3 [textbox] Name\n  [text] idle\n",
		v.Text,
	)
	v, e = s.Click("", "e1", true)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "[text] clicked", v.Text)
	assert.String(t, "", v.Landed)
	v, e = s.Fill(fill.New("", "e3", "Alfa"))
	assert.FatalOnError(t, e)
	assert.StringContains(t, "uid=e3 [textbox] Name = Alfa", v.Text)
	r := fill.New("", "e3", "Bravo")
	r.Direct = true
	v, e = s.Fill(r)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "uid=e3 [textbox] Name = Bravo", v.Text)
}

func TestClickLandsOnNextPage(t *testing.T) {
	b := browser_tester.New(t)
	s := b.Service()

	for uid, path := range map[string]string{
		"e1": "/next",
		"e2": "/slow",
		"e3": "/next",
	} {
		assert.FatalOnError(t, s.Navigate("", b.Site()))
		_, e := s.Snapshot("")
		assert.FatalOnError(t, e)
		v, f := s.Click("", uid, true)
		assert.FatalOnError(t, f)
		assert.Prefix(t, "[document] Next", v.Text)
		assert.Suffix(t, path, v.Landed)
	}
}

func TestScreenshotAfterSlowClick(t *testing.T) {
	b := browser_tester.New(t)
	s := b.Service()
	assert.FatalOnError(t, s.Navigate("", b.Site()))
	_, e := s.Snapshot("")
	assert.FatalOnError(t, e)
	_, e = s.Click("", "e2", false)
	assert.FatalOnError(t, e)
	_, e = s.Screenshot("")
	assert.FatalOnError(t, e)
}
