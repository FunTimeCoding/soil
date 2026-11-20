package tracker_source

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/download_result"
	"os"
	"path/filepath"
	"time"
)

func (s *Source) Refresh(x context.Context) error {
	t := s.tracker.Load()

	if t == nil {
		if cached := s.loadCache(); cached != nil {
			s.tracker.Store(cached)
			t = cached
			modified := cached.Modified()
			s.verified.Store(&modified)
		}
	}

	var since time.Time

	if t != nil {
		since = t.Modified()
	}

	if e := os.MkdirAll(filepath.Dir(s.cache), 0o750); e != nil {
		return fmt.Errorf("create tracker cache: %w", e)
	}

	f, e := os.CreateTemp(filepath.Dir(s.cache), "download-*")

	if e != nil {
		return fmt.Errorf("create tracker download: %w", e)
	}

	defer removeLeftover(f.Name())
	var result *download_result.Result
	g := s.guard.Attempt(
		x,
		func(y context.Context) error {
			var h error
			result, h = s.client.Download(y, since, f)

			return h
		},
	)

	if h := f.Close(); g == nil {
		g = h
	}

	if g != nil {
		return g
	}

	checked := s.clock()

	if result.Changed {
		if h := s.install(f.Name(), result.Modified); h != nil {
			return h
		}
	}

	s.verified.Store(&checked)

	return nil
}
