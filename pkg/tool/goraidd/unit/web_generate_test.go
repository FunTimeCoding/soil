package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/web_interface_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestGeneratingHandsEliteFilesToParser(t *testing.T) {
	p := service_tester.NewParser(t)
	v := base.Generate(t, p, memory.New())
	status, _ := web_interface_tester.Generate(
		t,
		v,
		"WvW/20261024-201500.zevtc",
	)
	assert.Integer(t, http.StatusSeeOther, status)
	assert.Integer(t, 1, p.Generations())
	assert.StringContains(
		t,
		"20261024-201500_detailed_wvw_kill.json",
		p.Generated[0][0],
	)
}

func TestDownParserAnswersWithoutReportAndFailsFast(t *testing.T) {
	p := service_tester.NewParser(t)
	p.Fail(outage_tester.RefusedDial())
	r := memory.New()
	v := base.Generate(t, p, r)
	status, body := web_interface_tester.Generate(t, v, "20261024-201500.zevtc")
	assert.Integer(t, http.StatusServiceUnavailable, status)
	assert.StringContains(t, "goraidparsed unreachable since", body)
	again, _ := web_interface_tester.Generate(t, v, "20261024-201500.zevtc")
	assert.Integer(t, http.StatusServiceUnavailable, again)
	assert.Integer(t, 1, p.GenerateCalls())
	assert.Count(t, 0, r.Events())
}

func TestMissingInputAnswersNotFoundWithoutReport(t *testing.T) {
	p := service_tester.NewParser(t)
	p.Fail(not_found.New("input file", "foxtrot.json"))
	r := memory.New()
	status, _ := web_interface_tester.Generate(
		t,
		base.Generate(t, p, r),
		"foxtrot.zevtc",
	)
	assert.Integer(t, http.StatusNotFound, status)
	assert.Count(t, 0, r.Events())
}

func TestFailedGenerationIsReported(t *testing.T) {
	p := service_tester.NewParser(t)
	p.Fail(unexpected.Format("generate report status: 500: golf"))
	r := memory.New()
	status, _ := web_interface_tester.Generate(
		t,
		base.Generate(t, p, r),
		"golf.zevtc",
	)
	assert.Integer(t, http.StatusInternalServerError, status)
	assert.Count(t, 1, r.Events())
}
