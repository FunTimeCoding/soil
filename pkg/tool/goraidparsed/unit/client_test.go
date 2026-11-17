package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/reported"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/mock_parser"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/service_tester"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateReturnsReport(t *testing.T) {
	c := base.New(t, mock_parser.New(2, nil)).Client()
	result, e := c.PostGenerate(
		t.Context(),
		[]string{service_tester.Input(t, "charlie.zevtc")},
		nil,
	)
	assert.FatalOnError(t, e)
	assert.True(t, strings.HasPrefix(filepath.Base(result.Path), "ONYX Log "))
	assert.Integer(t, 2, result.TidCount)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestMissingInputFileIsNotFound(t *testing.T) {
	c := base.New(t, mock_parser.New(1, nil)).Client()
	absent := filepath.Join(t.TempDir(), "delta.zevtc")
	_, e := c.PostGenerate(t.Context(), []string{absent}, nil)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "delta.zevtc", e.Error())
}

func TestNoFilesAreCallersMistake(t *testing.T) {
	c := base.New(t, mock_parser.New(1, nil)).Client()
	_, e := c.PostGenerate(t.Context(), []string{}, nil)
	assert.True(t, validation.Is(e))
}

func TestFailedGenerateIsReported(t *testing.T) {
	c := base.New(t, mock_parser.New(0, errors.New("echo parse failed"))).Client()
	_, e := c.PostGenerate(
		t.Context(),
		[]string{service_tester.Input(t, "echo.zevtc")},
		nil,
	)
	assert.True(t, reported.Is(e))
}
