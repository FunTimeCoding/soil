package unit

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/index"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/packager"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/signature"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/types/listing"
	"github.com/funtimecoding/soil/pkg/system/unit/index_tester"
	"github.com/funtimecoding/soil/pkg/system/unit/package_server_tester"
	"github.com/funtimecoding/soil/pkg/system/unit/packager_tester"
	"maps"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestIndexRead(t *testing.T) {
	directory := t.TempDir()
	path := index_tester.WriteIndex(directory)
	entries, e := index.Read(path)
	errors.PanicOnError(e)
	assert.Count(t, 2, entries)
	assert.String(t, "gohw", entries[0].Name)
	assert.String(t, "0.11.96-r1", entries[0].Version)
	assert.String(t, "x86_64", entries[0].Architecture)
	assert.String(t, "gobuild", entries[1].Name)
}

func TestIndexes(t *testing.T) {
	directory := t.TempDir()
	index_tester.WriteIndex(
		filepath.Join(
			directory,
			"rolling",
			"main",
			constant.AlpineArchitecture,
		),
	)
	listings, e := listing.Indexes(directory)
	errors.PanicOnError(e)
	assert.Count(t, 1, listings)
	assert.String(t, "rolling", listings[0].Version)
	assert.String(t, "main", listings[0].Repository)
	assert.String(t, "x86_64", listings[0].Architecture)
	assert.Count(t, 2, listings[0].Packages)
}

func TestSign(t *testing.T) {
	d := t.TempDir()
	unsignedPath := filepath.Join(d, "test-unsigned.apk")
	package_server_tester.CreateTestPackage(d, unsignedPath)
	privateKey := package_server_tester.GenerateRSAKey(2048)
	unsignedSegments := package_server_tester.SplitArchive(unsignedPath)

	if len(unsignedSegments) != 2 {
		t.Fatalf(
			"unsigned package should have 2 segments, got %d",
			len(unsignedSegments),
		)
	}

	signature.MustSignPackageWithKey(unsignedPath, privateKey, "test.rsa")
	signedSegments := package_server_tester.SplitArchive(unsignedPath)

	if len(signedSegments) != 3 {
		t.Fatalf(
			"signed package should have 3 segments, got %d",
			len(signedSegments),
		)
	}

	signatureSegment := signedSegments[0]
	signatureFiles := package_server_tester.ReadTarGz(signatureSegment)
	var signature []byte

	for name, content := range signatureFiles {
		if strings.HasPrefix(name, constant.AlpineSignaturePrefix) {
			signature = content

			break
		}
	}

	if signature == nil {
		t.Fatalf(
			"signature segment missing .SIGN.RSA.* file, got files: %v",
			slices.Collect(maps.Keys(signatureFiles)),
		)
	}

	controlSegment := signedSegments[1]
	hash := sha1.Sum(controlSegment)
	errors.PanicOnError(
		rsa.VerifyPKCS1v15(
			&privateKey.PublicKey,
			crypto.SHA1,
			hash[:],
			signature,
		),
	)
	controlFiles := package_server_tester.ReadTarGz(signedSegments[1])

	if _, okay := controlFiles[constant.AlpineMetadataFile]; !okay {
		t.Errorf("control segment missing .PKGINFO")
	}

	archiveFiles := package_server_tester.ReadTarGz(signedSegments[2])

	if _, okay := archiveFiles["usr/bin/test-binary"]; !okay {
		t.Errorf("data segment missing binary")
	}
}

func TestAlpinePackager(t *testing.T) {
	actual := packager.New("goexample", soil.DefaultVersion)
	packager_tester.StripTemporaryPrefix(actual)
	assert.Any(
		t,
		&packager.Packager{
			ExecutablePath:   "goexample",
			ExecutableName:   "goexample",
			PackageName:      "goexample-1.0.0-r1.apk",
			PackageVersion:   "1.0.0-r1",
			WorkDirectory:    "/gopackageapk-goexample",
			ControlDirectory: "/gopackageapk-goexample/control",
			ArchiveDirectory: "/gopackageapk-goexample/data",
			OutputFile:       "goexample-1.0.0-r1.apk",
		},
		actual,
	)
}

func TestCreate(t *testing.T) {
	d := t.TempDir()
	scriptPath := filepath.Join(d, "hello")
	system.WriteFile(
		scriptPath,
		[]byte(
			`#!/bin/sh
echo "Hello from Alpine package!"
`,
		),
		0755,
	)
	p := packager.New(scriptPath, soil.DefaultVersion)
	assert.String(t, "hello-1.0.0-r1.apk", p.PackageName)
	p.WorkDirectory = filepath.Join(d, "workspace")
	p.ControlDirectory = filepath.Join(
		p.WorkDirectory,
		constant.AlpineControlDirectory,
	)
	p.ArchiveDirectory = filepath.Join(p.WorkDirectory, "data")
	p.OutputFile = filepath.Join(d, p.PackageName)
	p.CreateWorkspace()
	p.CopyBinary()
	p.WritePKGINFO(p.CreateArchive())
	p.CreateControlTar()
	p.ConcatenateTars()
	apkPath := p.OutputFile
	assert.FileExists(t, apkPath)
	controlTarPath := filepath.Join(p.WorkDirectory, constant.AlpineControlFile)
	archivePath := filepath.Join(p.WorkDirectory, constant.AlpineArchiveFile)
	declaredHash := checkControlTar(t, controlTarPath)
	checkArchiveTar(t, archivePath, declaredHash)
	checkConcatenated(t, apkPath, controlTarPath, archivePath)
}
