package packager_tester

import (
	"archive/tar"
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
	"testing"
)

func ArchiveTar(
	t *testing.T,
	declaredHash string,
	archivePath string,
) {
	t.Helper()
	archiveSum := sha256.Sum256(system.ReadBytesUnsafe(archivePath))
	assert.String(t, declaredHash, hex.EncodeToString(archiveSum[:]))
	archiveFile := system.Open(archivePath)
	defer errors.PanicClose(archiveFile)
	gzr := system.GnuZipReader(archiveFile)
	defer errors.PanicClose(gzr)
	tr := tar.NewReader(gzr)
	foundBinary := false
	foundPAXHeader := false

	for {
		h, e := tr.Next()

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
}
