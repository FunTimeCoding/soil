package generation

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"os"
	"path/filepath"
	"time"
)

func Open(
	root string,
	name string,
	now time.Time,
	owns func(string) bool,
) string {
	result := filepath.Join(root, name)
	errors.PanicOnError(os.MkdirAll(result, 0755))
	refresh(result, now)
	Collect(root, name, now, constant.GenerationRetention, owns)

	return result
}
