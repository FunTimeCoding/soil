package unit

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"io"
	"strings"
	"testing"
)

func checkControlTar(t *testing.T, controlTarPath string) string {
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

		if h.Name == constant.AlpineMetadataFile {
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

	return declaredHash
}
