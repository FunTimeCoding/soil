package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"time"
)

func (c *Cache) fileState(identifier string) (time.Time, int64, bool) {
	i, e := os.Stat(
		filepath.Join(
			c.Base(),
			join.Empty(identifier, constant.NotationLogExtension),
		),
	)

	if e != nil {
		return time.Time{}, 0, false
	}

	return i.ModTime(), i.Size(), true
}
