package break_file

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/markup"
	"os"
)

func Load(path string) *File {
	b, e := os.ReadFile(path)
	errors.PanicOnError(e)
	result := &File{}
	markup.MustDecode(string(b), result)

	return result
}
