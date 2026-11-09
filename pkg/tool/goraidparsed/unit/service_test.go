package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/fake_parser"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/service_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerateMovesReportOutUnderItsDate(t *testing.T) {
	p := fake_parser.New(3, nil)
	s, output := service_tester.New(t, p)
	v, e := s.Generate(
		[]string{inputFile(t, "alfa.zevtc")},
		time.Date(2026, 1, 2, 20, 0, 0, 0, time.UTC),
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		filepath.Join(output, "ONYX Log 2026-01-02 20-00.html"),
		v.Path,
	)
	assert.Integer(t, 3, v.TidCount)
	assert.String(
		t,
		"report",
		system.ReadFile(filepath.Dir(v.Path), filepath.Base(v.Path)),
	)
	assert.True(t, p.Workdir() != "")
	assert.False(t, system.DirectoryExists(p.Workdir()))
}

func TestTwoReportsMinuteApartBothSurvive(t *testing.T) {
	s, output := service_tester.New(t, fake_parser.New(1, nil))
	first := time.Date(2026, 1, 2, 20, 0, 0, 0, time.UTC)
	_, e := s.Generate([]string{inputFile(t, "foxtrot.zevtc")}, first)
	assert.FatalOnError(t, e)
	_, f := s.Generate(
		[]string{inputFile(t, "golf.zevtc")},
		first.Add(time.Minute),
	)
	assert.FatalOnError(t, f)
	assert.True(
		t,
		system.FileExists(
			filepath.Join(output, "ONYX Log 2026-01-02 20-00.html"),
		),
	)
	assert.True(
		t,
		system.FileExists(
			filepath.Join(output, "ONYX Log 2026-01-02 20-01.html"),
		),
	)
}

func TestGenerateAnswersMissingInput(t *testing.T) {
	p := fake_parser.New(1, nil)
	s, _ := service_tester.New(t, p)
	_, e := s.Generate(
		[]string{filepath.Join(t.TempDir(), "absent.zevtc")},
		time.Now(),
	)
	assert.True(t, not_found.Is(e))
	assert.Integer(t, 0, p.Calls())
}

func TestGenerateRefusesNoFiles(t *testing.T) {
	p := fake_parser.New(1, nil)
	s, _ := service_tester.New(t, p)
	_, e := s.Generate(nil, time.Now())
	assert.True(t, validation.Is(e))
	assert.String(t, "no files to parse", e.Error())
	assert.Integer(t, 0, p.Calls())
}

func TestFailedParseNamesStep(t *testing.T) {
	failure := errors.New("bravo parse failed")
	s, _ := service_tester.New(t, fake_parser.New(0, failure))
	_, e := s.Generate([]string{inputFile(t, "bravo.zevtc")}, time.Now())
	assert.True(t, errors.Is(e, failure))
	assert.String(t, "parse: bravo parse failed", e.Error())
}
