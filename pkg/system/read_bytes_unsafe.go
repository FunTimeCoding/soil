package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ReadBytesUnsafe(path string) []byte {
	result, e := os.ReadFile(path)
	errors.PanicOnError(e)

	return result
}
