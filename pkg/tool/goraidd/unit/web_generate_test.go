package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"net"
	"net/http"
	"os"
	"syscall"
	"testing"
)

func TestGeneratingHandsTheEliteFilesToTheParser(t *testing.T) {
	p := service_tester.NewParser(t)
	v := mountedGenerate(t, p, memory.New())
	status, _ := postGenerate(t, v, "WvW/20261024-201500.zevtc")
	assert.Integer(t, http.StatusSeeOther, status)
	assert.Integer(t, 1, p.Generations())
	assert.StringContains(
		t,
		"20261024-201500_detailed_wvw_kill.json",
		p.Generated[0][0],
	)
}

func TestADownParserAnswersWithoutAReportAndFailsFast(t *testing.T) {
	p := service_tester.NewParser(t)
	p.Fail(
		&net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
		},
	)
	r := memory.New()
	v := mountedGenerate(t, p, r)
	status, body := postGenerate(t, v, "20261024-201500.zevtc")
	assert.Integer(t, http.StatusServiceUnavailable, status)
	assert.StringContains(t, "goraidparsed unreachable since", body)
	again, _ := postGenerate(t, v, "20261024-201500.zevtc")
	assert.Integer(t, http.StatusServiceUnavailable, again)
	assert.Integer(t, 1, p.GenerateCalls())
	assert.Count(t, 0, r.Events())
}

func TestAMissingInputAnswersNotFoundWithoutAReport(t *testing.T) {
	p := service_tester.NewParser(t)
	p.Fail(not_found.New("input file", "foxtrot.json"))
	r := memory.New()
	status, _ := postGenerate(t, mountedGenerate(t, p, r), "foxtrot.zevtc")
	assert.Integer(t, http.StatusNotFound, status)
	assert.Count(t, 0, r.Events())
}

func TestAFailedGenerationIsReported(t *testing.T) {
	p := service_tester.NewParser(t)
	p.Fail(unexpected.Format("generate report status: 500: golf"))
	r := memory.New()
	status, _ := postGenerate(t, mountedGenerate(t, p, r), "golf.zevtc")
	assert.Integer(t, http.StatusInternalServerError, status)
	assert.Count(t, 1, r.Events())
}
