package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestTheClientUploadsToTheDaemon(t *testing.T) {
	o := service_tester.New(t)
	c := mountedClient(t, o)
	assert.FatalOnError(
		t,
		c.PostPackage(
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
			filepath.Join(o.PackageRoot, "v3.21", "main", "x86_64", "a.apk"),
		),
	)
}

func TestTheClientSeesARefusedUploadAsValidation(t *testing.T) {
	c := mountedClient(t, service_tester.New(t))
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

func TestTheClientListsWhatTheDaemonIndexed(t *testing.T) {
	c := mountedClient(t, service_tester.New(t))
	result, e := c.GetPackages(t.Context(), &client.GetPackagesParams{})
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}

func TestTheClientReachesGoalpined(t *testing.T) {
	c := mountedClient(t, service_tester.New(t))
	assert.FatalOnError(t, c.Reachable(t.Context()))
}
