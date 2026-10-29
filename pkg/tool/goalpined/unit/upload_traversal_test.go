package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"net/http"
	"path/filepath"
	"testing"
)

func TestAnUploadCannotLeaveThePackageRoot(t *testing.T) {
	o := service_tester.New(t)
	v := mountedUpload(t, o.Service)
	body := service_tester.PackageBytes(t)
	escape := filepath.Join(filepath.Dir(o.PackageRoot), "etc")

	for _, path := range []string{
		"/api/packages/%2e%2e/main/etc/a.apk",
		"/api/packages/v3.21/main/%2e%2e/a.apk",
		"/api/packages/v3.21/%2e%2e/x86_64/a.apk",
		"/api/packages/v3.21/main/x86_64/%2e%2e%2fa.apk",
	} {
		assert.Integer(t, http.StatusBadRequest, postUpload(t, v, path, body))
	}

	assert.False(t, system.FileExists(escape))
	assert.Count(t, 0, o.Indexer.Directories)
}

func TestAnUploadIsStoredSignedAndIndexed(t *testing.T) {
	o := service_tester.New(t)
	v := mountedUpload(t, o.Service)
	assert.Integer(
		t,
		http.StatusCreated,
		postUpload(
			t,
			v,
			constant.FixtureUploadPath,
			service_tester.PackageBytes(t),
		),
	)
	directory := filepath.Join(o.PackageRoot, "v3.21", "main", "x86_64")
	assert.True(t, system.FileExists(filepath.Join(directory, "a.apk")))
	assert.Count(t, 1, o.Indexer.Directories)
	assert.String(t, directory, o.Indexer.Directories[0])
}

func TestAnUploadMustBeAnApk(t *testing.T) {
	o := service_tester.New(t)
	v := mountedUpload(t, o.Service)
	assert.Integer(
		t,
		http.StatusBadRequest,
		postUpload(
			t,
			v,
			"/api/packages/v3.21/main/x86_64/a.txt",
			service_tester.PackageBytes(t),
		),
	)
}
