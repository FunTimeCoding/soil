package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestClientUploadsToDaemon(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(
		t,
		s.Client().PostPackage(
			t.Context(),
			"v3.21",
			"main",
			"x86_64",
			"a.apk",
			bytes.NewReader(service_tester.PackageBytes(t)),
		),
	)
	assert.True(
		t,
		system.FileExists(
			filepath.Join(s.PackageRoot, "v3.21", "main", "x86_64", "a.apk"),
		),
	)
}

func TestClientSeesRefusedUploadAsValidation(t *testing.T) {
	c := base.New(t).Client()
	e := c.PostPackage(
		t.Context(),
		"v3.21",
		"main",
		"x86_64",
		"a.apk",
		bytes.NewReader([]byte("not a package")),
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "not an apk package", e.Error())
}

func TestClientListsWhatDaemonIndexed(t *testing.T) {
	c := base.New(t).Client()
	result, e := c.GetPackages(t.Context(), &client.GetPackagesParams{})
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}

func TestClientReachesGoalpined(t *testing.T) {
	c := base.New(t).Client()
	assert.FatalOnError(t, c.Reachable(t.Context()))
}
