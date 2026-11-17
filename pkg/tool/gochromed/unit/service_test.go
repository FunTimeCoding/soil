package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/fill"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/service_tester"
	"testing"
)

func TestCloseTabResolvesByTitle(t *testing.T) {
	o := service_tester.New(t)
	o.Source.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	o.Source.AddTab("bravo", "Bravo Wiki", "https://bravo.example")
	result, e := o.Service().CloseTab(selector.New("", "Bravo", ""))
	assert.FatalOnError(t, e)
	assert.String(t, "bravo", result.Identifier)
	tabs, f := o.Source.Tabs()
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(tabs))
}

func TestCloseTabAnswersUnknownTab(t *testing.T) {
	o := service_tester.New(t)
	o.Source.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	_, e := o.Service().CloseTab(selector.New("charlie", "", ""))
	assert.True(t, not_found.Is(e))
}

func TestClickBeforeSnapshotAnswersNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Source.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	_, e := o.Service().Click(selector.New("", "", ""), "e3", false)
	assert.True(t, not_found.Is(e))
}

func TestFillRequiresValue(t *testing.T) {
	_, e := service_tester.New(t).Service().Fill(
		fill.New(selector.New("", "", ""), "e5", ""),
	)
	assert.True(t, validation.Is(e))
}

func TestScreenshotSavesToDownloadDirectory(t *testing.T) {
	o := service_tester.New(t)
	o.Source.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	result, e := o.Service().Screenshot(selector.New("", "", ""))
	assert.FatalOnError(t, e)
	assert.FileExists(t, result.Path)
	assert.String(t, "Alfa Portal", result.Title)
}
