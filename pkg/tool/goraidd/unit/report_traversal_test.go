package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
	"path/filepath"
	"testing"
)

func TestDeletingAReportStaysInsideTheReportDirectory(t *testing.T) {
	base := t.TempDir()
	output := filepath.Join(base, "reports")
	system.EnsurePathExists(output)
	outside := filepath.Join(base, "secret.db")
	system.WriteFile(outside, []byte("keep"), 0o600)
	report := filepath.Join(output, "alfa.html")
	system.WriteFile(report, []byte("report"), 0o600)
	v := mountedReports(t, output)
	status := deleteReport(t, v, "..%2Fsecret.db")
	assert.True(t, system.FileExists(outside))
	assert.Integer(t, http.StatusNotFound, status)
	assert.Integer(t, http.StatusSeeOther, deleteReport(t, v, "alfa.html"))
	assert.False(t, system.FileExists(report))
}
