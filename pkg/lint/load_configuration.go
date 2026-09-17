package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/markup"
	"os"
)

func loadConfiguration(path string) *configuration {
	result := &configuration{}

	if path == "" {
		return result
	}

	b, e := os.ReadFile(path)
	errors.PanicOnError(e)
	errors.PanicOnError(markup.Decode(string(b), result))

	return result
}
