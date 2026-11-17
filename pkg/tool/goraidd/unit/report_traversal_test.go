package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"net/http"
	"path/filepath"
	"testing"
)

func TestDeletingReportStaysInsideReportDirectory(t *testing.T) {
	b := t.TempDir()
	output := filepath.Join(b, "reports")
	system.EnsurePathExists(output)
	outside := filepath.Join(b, "secret.db")
	system.WriteFile(outside, []byte("keep"), 0o600)
	report := filepath.Join(output, "alfa.html")
	system.WriteFile(report, []byte("report"), 0o600)
	s := base.New(t, service_tester.WithOutput(output))
	status := s.Delete("..%2Fsecret.db")
	assert.True(t, system.FileExists(outside))
	assert.Integer(t, http.StatusNotFound, status)
	assert.Integer(t, http.StatusSeeOther, s.Delete("alfa.html"))
	assert.False(t, system.FileExists(report))
}
