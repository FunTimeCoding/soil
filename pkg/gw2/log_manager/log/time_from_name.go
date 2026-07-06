package log

import (
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	library "github.com/funtimecoding/soil/pkg/time"
	"path/filepath"
	"time"
)

func timeFromName(fileName string) time.Time {
	base := filepath.Base(fileName)
	date, _ := key_value.Dot(base)

	return library.Parse("20060102-150405", date)
}
