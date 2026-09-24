package source

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"os"
)

func IsGeneratedFile(path string) bool {
	f, e := os.Open(path)

	if e != nil {
		return false
	}

	defer errors.LogClose(f)
	b := make([]byte, constant.GeneratedBytes)
	n, e := f.Read(b)

	if e != nil {
		return false
	}

	return IsGeneratedHeader(string(b[:n]))
}
