package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"net/http"
	"path/filepath"
	"testing"
)

func TestUploadCannotLeavePackageRoot(t *testing.T) {
	s := base.New(t)
	body := service_tester.PackageBytes(t)
	escape := filepath.Join(filepath.Dir(s.PackageRoot), "etc")

	for _, path := range []string{
		"/api/packages/%2e%2e/main/etc/a.apk",
		"/api/packages/v3.21/main/%2e%2e/a.apk",
		"/api/packages/v3.21/%2e%2e/x86_64/a.apk",
		"/api/packages/v3.21/main/x86_64/%2e%2e%2fa.apk",
	} {
		assert.Integer(t, http.StatusBadRequest, s.Post(path, body))
	}

	assert.False(t, system.FileExists(escape))
	assert.Count(t, 0, s.Indexer.Directories)
}

func TestUploadIsStoredSignedAndIndexed(t *testing.T) {
	s := base.New(t)
	assert.Integer(
		t,
		http.StatusCreated,
		s.Post(constant.FixtureUploadPath, service_tester.PackageBytes(t)),
	)
	directory := filepath.Join(s.PackageRoot, "v3.21", "main", "x86_64")
	assert.True(t, system.FileExists(filepath.Join(directory, "a.apk")))
	assert.Count(t, 1, s.Indexer.Directories)
	assert.String(t, directory, s.Indexer.Directories[0])
}

func TestUploadMustBeApk(t *testing.T) {
	assert.Integer(
		t,
		http.StatusBadRequest,
		base.New(t).Post(
			"/api/packages/v3.21/main/x86_64/a.txt",
			service_tester.PackageBytes(t),
		),
	)
}
