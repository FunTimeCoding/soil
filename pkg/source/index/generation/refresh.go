package generation

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"os"
	"path/filepath"
	"time"
)

func refresh(
	directory string,
	now time.Time,
) {
	marker := filepath.Join(directory, constant.GenerationMarker)

	if i, e := os.Stat(marker); e == nil &&
		now.Sub(i.ModTime()) < constant.GenerationRefresh {
		return
	}

	errors.LogOnError(os.WriteFile(marker, nil, 0644))
	errors.LogOnError(os.Chtimes(marker, now, now))
}
