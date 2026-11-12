package generation

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"os"
	"path/filepath"
	"time"
)

func touched(
	directory string,
	now time.Time,
) time.Time {
	if i, e := os.Stat(
		filepath.Join(directory, constant.GenerationMarker),
	); e == nil {
		return i.ModTime()
	}

	if i, e := os.Stat(directory); e == nil {
		return i.ModTime()
	}

	return now
}
