package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/markup"
	"io/fs"
	"os"
	"path/filepath"
)

func loadConfiguration(root string) *configuration {
	result := &configuration{}
	b, e := os.ReadFile(filepath.Join(root, constant.ConfigurationPath))

	if errors.Is(e, fs.ErrNotExist) {
		return result
	}

	errors.PanicOnError(e)
	errors.PanicOnError(markup.Decode(string(b), result))

	return result
}
