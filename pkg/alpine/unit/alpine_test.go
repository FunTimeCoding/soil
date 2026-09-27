package unit

import (
	"archive/tar"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/alpine/packager"
	"github.com/funtimecoding/soil/pkg/alpine/unit/index_tester"
	"github.com/funtimecoding/soil/pkg/alpine/unit/package_server_tester"
	"github.com/funtimecoding/soil/pkg/alpine/unit/packager_tester"
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
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
		filepath.Join(directory, "rolling", "main", constant.Architecture),
	)
	listings, e := package_server.Indexes(directory)
	errors.PanicOnError(e)
	assert.Count(t, 1, listings)
	assert.String(t, "rolling", listings[0].Version)
	assert.String(t, "main", listings[0].Repository)
	assert.String(t, "x86_64", listings[0].Architecture)
	assert.Count(t, 2, listings[0].Packages)
}

func TestSign(t *testing.T) {
	d := filepath.Join("../../..", "tmp", "alpine-signing-test")
	system.MakeDirectory(d)
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

	package_server.SignPackageWithKey(unsignedPath, privateKey, "test.rsa")
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
		if strings.HasPrefix(name, constant.SignaturePrefix) {
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

	if _, okay := controlFiles[constant.MetadataFile]; !okay {
		t.Errorf("control segment missing .PKGINFO")
	}

	archiveFiles := package_server_tester.ReadTarGz(signedSegments[2])

	if _, okay := archiveFiles["usr/bin/test-binary"]; !okay {
		t.Errorf("data segment missing binary")
	}
}

func TestPackager(t *testing.T) {
	actual := packager.New("goexample", library.DefaultVersion)
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
	d := filepath.Join("../../..", "tmp", "alpine-package")
	system.MakeDirectory(d)
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
	p := packager.New(scriptPath, library.DefaultVersion)
	assert.String(t, "hello-1.0.0-r1.apk", p.PackageName)
	p.WorkDirectory = filepath.Join(d, "workspace")
	p.ControlDirectory = filepath.Join(p.WorkDirectory, "control")
	p.ArchiveDirectory = filepath.Join(p.WorkDirectory, "data")
	p.OutputFile = filepath.Join(d, p.PackageName)
	p.CreateWorkspace()
	p.CopyBinary()
	p.WritePKGINFO(p.CreateArchive())
	p.CreateControlTar()
	p.ConcatenateTars()
	apkPath := p.OutputFile
	assert.FileExists(t, apkPath)
	controlTarPath := filepath.Join(p.WorkDirectory, constant.ControlFile)
	archivePath := filepath.Join(p.WorkDirectory, constant.ArchiveFile)
	controlFile := system.Open(controlTarPath)
	defer errors.PanicClose(controlFile)
	gzr := system.GnuZipReader(controlFile)
	defer errors.PanicClose(gzr)
	tr := tar.NewReader(gzr)
	foundPackageInformation := false
	declaredHash := ""

	for {
		h, e := tr.Next()

		if e == io.EOF {
			break
		}

		errors.PanicOnError(e)

		if h.Name == constant.MetadataFile {
			foundPackageInformation = true
			content := string(system.ReadAll(tr))
			assert.StringContains(t, "pkgname = hello", content)
			assert.StringContains(t, "pkgver = 1.0.0-r1", content)
			assert.StringContains(t, "size = ", content)
			assert.StringContains(t, "datahash =", content)

			for _, line := range strings.Split(content, "\n") {
				if value, found := strings.CutPrefix(
					line,
					"datahash = ",
				); found {
					declaredHash = value
				}
			}
		}
	}

	if !foundPackageInformation {
		t.Errorf("control.tar.gz missing .PKGINFO")
	}

	archiveSum := sha256.Sum256(system.ReadBytesUnsafe(archivePath))
	assert.String(t, hex.EncodeToString(archiveSum[:]), declaredHash)
	archiveFile := system.Open(archivePath)
	defer errors.PanicClose(archiveFile)
	gzr2 := system.GnuZipReader(archiveFile)
	defer errors.PanicClose(gzr2)
	tr2 := tar.NewReader(gzr2)
	foundBinary := false
	foundPAXHeader := false

	for {
		h, e := tr2.Next()

		if e == io.EOF {
			break
		}

		errors.PanicOnError(e)

		if h.Name == "usr/bin/hello" {
			foundBinary = true

			if !h.FileInfo().Mode().IsRegular() {
				t.Errorf("binary is not a regular file")
			}

			if _, okay := h.PAXRecords["APK-TOOLS.checksum.SHA1"]; okay {
				foundPAXHeader = true
			}
		}
	}

	if !foundBinary {
		t.Errorf("data.tar.gz missing usr/bin/hello")
	}

	if !foundPAXHeader {
		t.Errorf("data.tar.gz missing PAX header with SHA1 checksum")
	}

	apkStat := system.Stat(apkPath)
	controlStat := system.Stat(controlTarPath)
	archiveStat := system.Stat(archivePath)
	expectedSize := controlStat.Size() + archiveStat.Size()

	if apkStat.Size() != expectedSize {
		t.Errorf(
			"APK size mismatch: got %d, expected %d (control %d + data %d)",
			apkStat.Size(),
			expectedSize,
			controlStat.Size(),
			archiveStat.Size(),
		)
	}

	apkFile := system.Open(apkPath)
	defer errors.PanicClose(apkFile)
	controlBytes := make([]byte, controlStat.Size())
	system.ReadFull(apkFile, controlBytes)
	controlGz := system.GnuZipReader(strings.NewReader(string(controlBytes)))
	errors.PanicClose(controlGz)
	archiveBytes := make([]byte, archiveStat.Size())
	system.ReadFull(apkFile, archiveBytes)
	archiveGz := system.GnuZipReader(strings.NewReader(string(archiveBytes)))
	defer errors.PanicClose(archiveGz)
	archive := tar.NewReader(archiveGz)

	for {
		h, e := archive.Next()

		if e == io.EOF {
			break
		}

		errors.PanicOnError(e)

		if h.Name == "usr/bin/hello" {
			content := string(system.ReadAll(archive))

			if !strings.Contains(content, "#!/bin/sh") {
				t.Errorf("binary missing shebang")
			}

			if !strings.Contains(
				content,
				"Hello from Alpine package!",
			) {
				t.Errorf("binary missing expected message")
			}

			if _, okay := h.PAXRecords["APK-TOOLS.checksum.SHA1"]; !okay {
				t.Errorf("missing SHA1 checksum in PAX headers")
			}
		}
	}
}
