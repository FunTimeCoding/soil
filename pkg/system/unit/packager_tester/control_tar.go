package packager_tester

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"io"
	"strings"
)

func (o *Tester) ControlTar(controlTarPath string) string {
	o.t.Helper()
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
			assert.StringContains(o.t,"pkgname = hello", content)
			assert.StringContains(o.t,"pkgver = 1.0.0-r1", content)
			assert.StringContains(o.t,"size = ", content)
			assert.StringContains(o.t,"datahash =", content)

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
		o.t.Errorf("control.tar.gz missing .PKGINFO")
	}

	return declaredHash
}
