package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestGeneratingHandsEliteFilesToParser(t *testing.T) {
	s := base.New(t)
	status, _ := s.Generate("WvW/20261024-201500.zevtc")
	assert.Integer(t, http.StatusSeeOther, status)
	assert.Integer(t, 1, s.Parser.Generations())
	assert.StringContains(
		t,
		"20261024-201500_detailed_wvw_kill.json",
		s.Parser.Generated[0][0],
	)
}

func TestDownParserAnswersWithoutReportAndFailsFast(t *testing.T) {
	s := base.New(t)
	s.Parser.Fail(outage_tester.RefusedDial())
	status, body := s.Generate("20261024-201500.zevtc")
	assert.Integer(t, http.StatusServiceUnavailable, status)
	assert.StringContains(t, "goraidparsed unreachable since", body)
	again, _ := s.Generate("20261024-201500.zevtc")
	assert.Integer(t, http.StatusServiceUnavailable, again)
	assert.Integer(t, 1, s.Parser.GenerateCalls())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestMissingInputAnswersNotFoundWithoutReport(t *testing.T) {
	s := base.New(t)
	s.Parser.Fail(not_found.New("input file", "foxtrot.json"))
	status, _ := s.Generate("foxtrot.zevtc")
	assert.Integer(t, http.StatusNotFound, status)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestFailedGenerationIsReported(t *testing.T) {
	s := base.New(t)
	s.Parser.Fail(unexpected.Format("generate report status: 500: golf"))
	status, _ := s.Generate("golf.zevtc")
	assert.Integer(t, http.StatusInternalServerError, status)
	assert.Count(t, 1, s.Reporter.Events())
}
