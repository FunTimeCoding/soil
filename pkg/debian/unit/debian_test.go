package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/debian/aptly"
	"github.com/funtimecoding/soil/pkg/debian/aptly/face"
	"github.com/funtimecoding/soil/pkg/debian/aptly/mock_client"
	debianConstant "github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/debian/packager"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
	"testing"
)

func TestAptlyClientSatisfiesTheRepositoryFace(t *testing.T) {
	var v face.Repository = aptly.New(
		"host.example",
		443,
		false,
		"user",
		"password",
	)
	assert.NotNil(t, v)
}

func TestAptlyMockSatisfiesTheRepositoryFace(t *testing.T) {
	var v face.Repository = mock_client.New()
	assert.NotNil(t, v)
}

func TestAptlyMockServesSeededVersions(t *testing.T) {
	c := mock_client.New()
	c.SeedVersions("stable", "foxtrot", "1.2.5", "1.2.3", "1.2.0")
	v, e := c.Versions("stable", "foxtrot")
	assert.Nil(t, e)
	assert.Count(t, 3, v)
	assert.String(t, "1.2.5", v[0])
}

func TestAptlyMockLatestVersionTakesTheFirst(t *testing.T) {
	c := mock_client.New()
	c.SeedVersions("stable", "foxtrot", "1.2.5", "1.2.3")
	v, e := c.LatestVersion("stable", "foxtrot")
	assert.Nil(t, e)
	assert.String(t, "1.2.5", v)
}

func TestAptlyMockReportsAnAbsentPackage(t *testing.T) {
	c := mock_client.New()
	v, e := c.LatestVersion("stable", "absent")
	assert.Nil(t, e)
	assert.String(t, "", v)
}

func TestAptlyMockPropagatesFailure(t *testing.T) {
	c := mock_client.New()
	c.Fail(errors.New("repository unreachable"))
	_, e := c.Versions("stable", "foxtrot")
	assert.NotNil(t, e)
}

func TestPackageVersionsReadsTheAptlyKeyFormat(t *testing.T) {
	v := aptly.PackageVersions(packageKeys(), "foxtrot")
	assert.Count(t, 3, v)
	assert.String(t, "1.2.1", v[0])
	assert.String(t, "1.2.5", v[1])
}

func TestPackageVersionsRefusesAnotherPackage(t *testing.T) {
	assert.Count(t, 2, aptly.PackageVersions(packageKeys(), "gohw"))
}

func TestPackageVersionsRefusesAnAbsentPackage(t *testing.T) {
	assert.Count(t, 0, aptly.PackageVersions(packageKeys(), "gooutpostd"))
}

func TestPackageVersion(t *testing.T) {
	assert.String(
		t,
		"example_1.0.0-1_amd64",
		debian.PackageVersion(
			"example",
			library.DefaultVersion,
			1,
			systemConstant.AMD64,
		),
	)
}

func TestRenderControl(t *testing.T) {
	assert.String(
		t,
		`Package: goexample
Version: 1.0.0
Architecture: amd64
Maintainer: John Doe <john.doe@example.org>
Description: Short stub description.
 Long stub description.
`,
		debian.RenderControl(
			"goexample",
			systemConstant.AMD64,
			library.DefaultVersion,
			"John Doe",
			"john.doe@example.org",
		),
	)
}

func TestRenderPostInstallRestart(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "configure" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
systemctl restart Alfa.service
`,
		debian.RenderPostInstall(
			constant.UpperAlfa,
			debianConstant.UpgradeRestart,
		),
	)
}

func TestRenderPostInstallKeep(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "configure" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
`,
		debian.RenderPostInstall(
			constant.UpperAlfa,
			debianConstant.UpgradeKeep,
		),
	)
}

func TestRenderPreRemove(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "remove" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl stop Alfa.service
systemctl disable Alfa.service
`,
		debian.RenderPreRemove(constant.UpperAlfa),
	)
}

func TestRenderPostRemove(t *testing.T) {
	assert.String(
		t,
		`#!/bin/sh
set -e

if [ "$1" != "remove" ] && [ "$1" != "purge" ]; then
    exit 0
fi

if [ ! -d /run/systemd/system ]; then
    exit 0
fi

systemctl daemon-reload
`,
		debian.RenderPostRemove(),
	)
}

func TestRenderUnit(t *testing.T) {
	assert.String(
		t,
		`[Unit]
Description=Alfa stub description
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/Alfa

[Install]
WantedBy=multi-user.target
`,
		debian.RenderUnit(constant.UpperAlfa),
	)
}

func TestPackager(t *testing.T) {
	actual := packager.New(
		"goexample",
		library.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	root := packageRoot()
	assert.Any(
		t,
		&packager.Packager{
			ExecutablePath:    "goexample",
			ExecutableName:    "goexample",
			PackageName:       "goexample_1.0.0-1_amd64",
			PackageVersion:    "1.0.0",
			Root:              root,
			ConfigurationRoot: filepath.Join(root, "DEBIAN"),
			UnitRoot: filepath.Join(
				root,
				"usr",
				"lib",
				"systemd",
				"system",
			),
			ControlFile:    filepath.Join(root, "DEBIAN", "control"),
			BinaryRoot:     filepath.Join(root, "usr", "local", "bin"),
			Architecture:   "amd64",
			MaintainerName: "John Doe",
			MaintainerMail: "john.doe@example.org",
		},
		actual,
	)
}

func TestSubPath(t *testing.T) {
	actual := packager.New(
		"tmp/goexample",
		library.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	root := packageRoot()
	assert.Any(
		t,
		&packager.Packager{
			ExecutablePath:    "tmp/goexample",
			ExecutableName:    "goexample",
			PackageName:       "goexample_1.0.0-1_amd64",
			PackageVersion:    "1.0.0",
			Root:              root,
			ConfigurationRoot: filepath.Join(root, "DEBIAN"),
			UnitRoot: filepath.Join(
				root,
				"usr",
				"lib",
				"systemd",
				"system",
			),
			ControlFile:    filepath.Join(root, "DEBIAN", "control"),
			BinaryRoot:     filepath.Join(root, "usr", "local", "bin"),
			Architecture:   "amd64",
			MaintainerName: "John Doe",
			MaintainerMail: "john.doe@example.org",
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
		library.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	p.BinaryRoot = filepath.Join(directory, "usr", "local", "bin")
	p.MoveBinary()
	assert.True(t, system.FileExists(filepath.Join(p.BinaryRoot, "goexample")))
}
