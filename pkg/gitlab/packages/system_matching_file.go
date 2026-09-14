package packages

import (
	"gitlab.com/gitlab-org/api/client-go/v3"
	"runtime"
	"strings"
)

func SystemMatchingFile(v []*gitlab.PackageFile) *gitlab.PackageFile {
	for _, e := range v {
		if !strings.Contains(e.FileName, runtime.GOARCH) {
			continue
		}

		if !strings.Contains(e.FileName, runtime.GOOS) {
			continue
		}

		return e
	}

	return nil
}
