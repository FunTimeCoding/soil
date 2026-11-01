package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly/face"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly/mock_client"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/packager"
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
	defer c.Close()
	c.SeedVersions("stable", "foxtrot", "1.2.5", "1.2.3", "1.2.0")
	v, e := c.Versions("stable", "foxtrot")
	assert.Nil(t, e)
	assert.Count(t, 3, v)
	assert.String(t, "1.2.5", v[0])
}

func TestAptlyMockLatestVersionTakesTheFirst(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.SeedVersions("stable", "foxtrot", "1.2.5", "1.2.3")
	v, e := c.LatestVersion("stable", "foxtrot")
	assert.Nil(t, e)
	assert.String(t, "1.2.5", v)
}

func TestAptlyMockReportsAnAbsentPackage(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	v, e := c.LatestVersion("stable", "absent")
	assert.Nil(t, e)
	assert.String(t, "", v)
}

func TestAptlyMockPropagatesFailure(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
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
			soil.DefaultVersion,
			1,
			constant.AMD64,
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
			constant.AMD64,
			soil.DefaultVersion,
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
			stringsConstant.UpperAlfa,
			constant.DebianUpgradeRestart,
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
			stringsConstant.UpperAlfa,
			constant.DebianUpgradeKeep,
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
		debian.RenderPreRemove(stringsConstant.UpperAlfa),
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
		debian.RenderUnit(stringsConstant.UpperAlfa),
	)
}

func TestDebianPackager(t *testing.T) {
	actual := packager.New(
		"goexample",
		soil.DefaultVersion,
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
		soil.DefaultVersion,
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
		soil.DefaultVersion,
		"John Doe",
		"john.doe@example.org",
	)
	p.BinaryRoot = filepath.Join(
		directory,
		constant.Resources,
		constant.Local,
		constant.Binary,
	)
	p.MoveBinary()
	assert.True(t, system.FileExists(filepath.Join(p.BinaryRoot, "goexample")))
}
