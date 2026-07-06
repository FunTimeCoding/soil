package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func WriteFile(
	name string,
	b []byte,
	m os.FileMode,
) {
	errors.PanicOnError(os.WriteFile(name, b, m))
}
