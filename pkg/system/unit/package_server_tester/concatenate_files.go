package package_server_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ConcatenateFiles(
	output string,
	files ...string,
) {
	f, e := os.Create(output)
	errors.PanicOnError(e)
	defer errors.PanicClose(f)

	for _, i := range files {
		b, g := os.ReadFile(i)
		errors.PanicOnError(g)
		_, h := f.Write(b)
		errors.PanicOnError(h)
	}
}
