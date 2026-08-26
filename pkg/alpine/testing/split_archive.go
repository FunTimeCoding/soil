package testing

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func SplitArchive(apkPath string) [][]byte {
	result, e := os.ReadFile(apkPath)
	errors.PanicOnError(e)

	return SplitArchiveBytes(result)
}
