package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/debian/packager"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestPackager(t *testing.T) {
	actual := packager.New(
		"goexample",
		constant.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	generifyPackager(actual)
	assert.Any(
		t,
		&packager.Packager{
			ExecutablePath:    "goexample",
			ExecutableName:    "goexample",
			PackageName:       "goexample_1.0.0-1_amd64",
			PackageVersion:    "1.0.0",
			Root:              "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64",
			ConfigurationRoot: "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/DEBIAN",
			UnitRoot:          "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/lib/systemd/system",
			ControlFile:       "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/DEBIAN/control",
			BinaryRoot:        "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/usr/local/bin",
			Architecture:      "amd64",
			MaintainerName:    "John Doe",
			MaintainerMail:    "john.doe@example.org",
		},
		actual,
	)
}

func TestSubPath(t *testing.T) {
	actual := packager.New(
		"tmp/goexample",
		constant.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	generifyPackager(actual)
	assert.Any(
		t,
		&packager.Packager{
			ExecutablePath:    "tmp/goexample",
			ExecutableName:    "goexample",
			PackageName:       "goexample_1.0.0-1_amd64",
			PackageVersion:    "1.0.0",
			Root:              "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64",
			ConfigurationRoot: "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/DEBIAN",
			UnitRoot:          "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/lib/systemd/system",
			ControlFile:       "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/DEBIAN/control",
			BinaryRoot:        "/soil/pkg/debian/unit_test/goexample_1.0.0-1_amd64/usr/local/bin",
			Architecture:      "amd64",
			MaintainerName:    "John Doe",
			MaintainerMail:    "john.doe@example.org",
		},
		actual,
	)
}

func TestMoveBinaryAbsolutePath(t *testing.T) {
	directory := t.TempDir()
	source := filepath.Join(directory, "goexample")
	system.WriteFile(source, []byte("binary"), 0755)
	p := packager.New(
		source,
		constant.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	p.BinaryRoot = filepath.Join(directory, "usr", "local", "bin")
	p.MoveBinary()
	assert.True(t, system.FileExists(filepath.Join(p.BinaryRoot, "goexample")))
}

func generifyPackager(p *packager.Packager) {
	d := "soil"
	p.Root = system.StripUntilDirectory(p.Root, d)
	p.ConfigurationRoot = system.StripUntilDirectory(p.ConfigurationRoot, d)
	p.UnitRoot = system.StripUntilDirectory(p.UnitRoot, d)
	p.ControlFile = system.StripUntilDirectory(p.ControlFile, d)
	p.BinaryRoot = system.StripUntilDirectory(p.BinaryRoot, d)
}
