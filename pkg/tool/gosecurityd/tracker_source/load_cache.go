package tracker_source

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
	"os"
	"time"
)

func (s *Source) loadCache() *tracker.Tracker {
	f, e := os.Open(s.cache)

	if e != nil {
		return nil
	}

	defer errors.LogClose(f)
	i, g := f.Stat()

	if g != nil {
		return nil
	}

	var modified time.Time

	if m := i.ModTime(); m.Unix() != 0 {
		modified = m.UTC()
	}

	result, h := tracker.Read(f, modified, s.releases...)

	if h != nil {
		return nil
	}

	return result
}
