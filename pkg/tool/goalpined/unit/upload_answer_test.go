package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestBodyThatIsNotPackageIsRefusedQuietly(t *testing.T) {
	s := base.New(t)
	assert.Integer(
		t,
		http.StatusBadRequest,
		s.Post(constant.FixtureUploadPath, []byte("not a package")),
	)
	assert.Count(t, 0, s.Reporter.Events())
	assert.Count(t, 0, s.Indexer.Directories)
	assert.False(
		t,
		system.FileExists(
			filepath.Join(s.PackageRoot, "v3.21", "main", "x86_64", "a.apk"),
		),
	)
}

func TestMissingSignatureKeyIsCaptured(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(
		t,
		os.Remove(filepath.Join(s.KeyDirectory, s.SignatureKey)),
	)
	assert.Integer(
		t,
		http.StatusInternalServerError,
		s.Post(constant.FixtureUploadPath, service_tester.PackageBytes(t)),
	)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestFailedIndexIsCaptured(t *testing.T) {
	s := base.New(t)
	s.Indexer.Failure = errors.Format("index packages", "apk")
	assert.Integer(
		t,
		http.StatusInternalServerError,
		s.Post(constant.FixtureUploadPath, service_tester.PackageBytes(t)),
	)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestEmptyRootListsNothing(t *testing.T) {
	o := service_tester.New(t)
	assert.False(t, system.FileExists(o.PackageRoot))
	result, e := o.Service.Packages("")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}
