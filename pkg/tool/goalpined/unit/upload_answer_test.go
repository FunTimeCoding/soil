package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestABodyThatIsNotAPackageIsRefusedQuietly(t *testing.T) {
	o := service_tester.New(t)
	r := memory.New()
	v := mountedReported(t, o.Service, r)
	assert.Integer(
		t,
		http.StatusBadRequest,
		postUpload(t, v, constant.FixtureUploadPath, []byte("not a package")),
	)
	assert.Count(t, 0, r.Events())
	assert.Count(t, 0, o.Indexer.Directories)
	assert.False(
		t,
		system.FileExists(
			filepath.Join(o.PackageRoot, "v3.21", "main", "x86_64", "a.apk"),
		),
	)
}

func TestAMissingSignatureKeyIsCaptured(t *testing.T) {
	o := service_tester.New(t)
	assert.FatalOnError(
		t,
		os.Remove(filepath.Join(o.KeyDirectory, o.SignatureKey)),
	)
	r := memory.New()
	v := mountedReported(t, o.Service, r)
	assert.Integer(
		t,
		http.StatusInternalServerError,
		postUpload(
			t,
			v,
			constant.FixtureUploadPath,
			service_tester.PackageBytes(t),
		),
	)
	assert.Count(t, 1, r.Events())
}

func TestAFailedIndexIsCaptured(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = errors.Format("index packages", "apk")
	r := memory.New()
	v := mountedReported(t, o.Service, r)
	assert.Integer(
		t,
		http.StatusInternalServerError,
		postUpload(
			t,
			v,
			constant.FixtureUploadPath,
			service_tester.PackageBytes(t),
		),
	)
	assert.Count(t, 1, r.Events())
}

func TestAnEmptyRootListsNothing(t *testing.T) {
	o := service_tester.New(t)
	assert.False(t, system.FileExists(o.PackageRoot))
	result, e := o.Service.Packages("")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}
