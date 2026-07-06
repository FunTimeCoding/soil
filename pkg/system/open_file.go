package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
	"os"
)

func OpenFile(
	name string,
	m fs.FileMode,
) *os.File {
	result, e := os.OpenFile(
		name,
		os.O_CREATE|os.O_RDWR,
		m,
	)
	errors.PanicOnError(e)

	return result
}
