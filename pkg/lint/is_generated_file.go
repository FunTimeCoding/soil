package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func IsGeneratedFile(path string) bool {
	f, e := os.Open(path)

	if e != nil {
		return false
	}

	defer errors.LogClose(f)
	b := make([]byte, 512)
	n, e := f.Read(b)

	if e != nil {
		return false
	}

	return IsGeneratedHeader(string(b[:n]))
}
