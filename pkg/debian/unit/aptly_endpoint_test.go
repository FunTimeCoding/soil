package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackagesReadsTheList(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`["Pamd64 charlie 1.2.3 abc"]`,
	)
	result, e := newAptlyClient(t, s.URL).Packages("delta")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"Pamd64 charlie 1.2.3 abc"}, result)
	r := last()
	assert.String(t, "/api/repos/delta/packages", r.URL.Path)
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo", password)
}

func TestRefusedPackagesIsAnErrorNotAPanic(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"error":"authorization required"}`,
	)
	_, e := newAptlyClient(t, s.URL).Packages("delta")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "authorization required", e.Error())
}

func TestUploadSendsMultipart(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `["echo.deb"]`)
	path := filepath.Join(t.TempDir(), "echo.deb")
	assert.FatalOnError(t, os.WriteFile(path, []byte("foxtrot"), 0o600))
	assert.FatalOnError(t, newAptlyClient(t, s.URL).Upload("golf", path))
	r := last()
	assert.String(t, "/api/files/golf", r.URL.Path)
	assert.True(
		t,
		strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data"),
	)
}

func TestMissingUploadFileIsAnError(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusOK, "")
	e := newAptlyClient(t, s.URL).Upload("golf", "/nonexistent/hotel.deb")
	assert.True(t, e != nil)
}
