package unit

import (
	"archive/tar"
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/packager"
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackager(t *testing.T) {
	actual := packager.New("goexample", library.DefaultVersion)
	generifyPackager(actual)
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

func generifyPackager(p *packager.Packager) {
	d := "gopackageapk-goexample"
	p.WorkDirectory = system.StripUntilDirectory(p.WorkDirectory, d)
	p.ControlDirectory = system.StripUntilDirectory(p.ControlDirectory, d)
	p.ArchiveDirectory = system.StripUntilDirectory(p.ArchiveDirectory, d)
}
