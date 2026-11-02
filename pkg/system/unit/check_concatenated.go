package unit

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
	"strings"
	"testing"
)

func checkConcatenated(t *testing.T, apkPath string, controlTarPath string, archivePath string) {
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
