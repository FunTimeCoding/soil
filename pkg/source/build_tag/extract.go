package build_tag

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Extract(path string) []string {
	file, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(file)

	return extractFrom(file)
}
