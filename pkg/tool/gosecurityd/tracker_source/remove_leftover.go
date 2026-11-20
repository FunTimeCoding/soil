package tracker_source

import (
	"errors"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
	"os"
)

func removeLeftover(path string) {
	if e := os.Remove(path); !errors.Is(e, fs.ErrNotExist) {
		soilErrors.LogOnError(e)
	}
}
