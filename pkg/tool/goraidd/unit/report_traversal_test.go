package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/web_interface_tester"
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
	v := base.Reports(t, output)
	status := web_interface_tester.Delete(t, v, "..%2Fsecret.db")
	assert.True(t, system.FileExists(outside))
	assert.Integer(t, http.StatusNotFound, status)
	assert.Integer(
		t,
		http.StatusSeeOther,
		web_interface_tester.Delete(t, v, "alfa.html"),
	)
	assert.False(t, system.FileExists(report))
}
