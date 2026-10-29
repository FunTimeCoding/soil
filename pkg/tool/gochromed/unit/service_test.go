package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/fill"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/service_tester"
	"testing"
)

func TestCloseTabResolvesByTitle(t *testing.T) {
	c := mock_client.New()
	c.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	c.AddTab("bravo", "Bravo Wiki", "https://bravo.example")
	result, e := service_tester.New(t, c).CloseTab(
		selector.New("", "Bravo", ""),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo", result.Identifier)
	tabs, f := c.Tabs()
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(tabs))
}

func TestCloseTabAnswersAnUnknownTab(t *testing.T) {
	c := mock_client.New()
	c.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	_, e := service_tester.New(t, c).CloseTab(selector.New("charlie", "", ""))
	assert.True(t, not_found.Is(e))
}

func TestClickBeforeSnapshotAnswersNotFound(t *testing.T) {
	c := mock_client.New()
	c.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	_, e := service_tester.New(t, c).Click(
		selector.New("", "", ""),
		"e3",
		false,
	)
	assert.True(t, not_found.Is(e))
}

func TestFillRequiresAValue(t *testing.T) {
	_, e := service_tester.New(t, mock_client.New()).Fill(
		fill.New(selector.New("", "", ""), "e5", ""),
	)
	assert.True(t, validation.Is(e))
}

func TestScreenshotSavesToTheDownloadDirectory(t *testing.T) {
	c := mock_client.New()
	c.AddTab("alfa", "Alfa Portal", "https://alfa.example")
	result, e := service_tester.New(t, c).Screenshot(selector.New("", "", ""))
	assert.FatalOnError(t, e)
	assert.FileExists(t, result.Path)
	assert.String(t, "Alfa Portal", result.Title)
}
