package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"net/http"
	"path/filepath"
	"testing"
)

func TestReportsListOnlyGeneratedPages(t *testing.T) {
	o := service_tester.New(t)
	system.WriteFile(filepath.Join(o.Output, "alfa.html"), []byte("a"), 0o600)
	system.WriteFile(filepath.Join(o.Output, "bravo.txt"), []byte("b"), 0o600)
	system.EnsurePathExists(filepath.Join(o.Output, "charlie.html"))
	s := o.Service()
	reports, e := s.Reports()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, reports)
	assert.String(t, "alfa.html", reports[0].Name)
}

func TestDeletingMissingReportIsNotFound(t *testing.T) {
	s := service_tester.New(t).Service()
	assert.True(t, not_found.Is(s.DeleteReport("delta.html")))
}

func TestDeletingMissingReportAnswersNotFound(t *testing.T) {
	assert.Integer(t, http.StatusNotFound, base.New(t).Delete("echo.html"))
}

func TestGeneratingReadsEliteFileOfEachLog(t *testing.T) {
	s := service_tester.New(t).Service()
	files := s.EliteFiles([]string{"WvW/20261024-201500.zevtc"})
	assert.Count(t, 1, files)
	assert.StringContains(t, "20261024-201500_detailed_wvw_kill.json", files[0])
}
