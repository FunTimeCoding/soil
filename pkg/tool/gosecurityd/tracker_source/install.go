package tracker_source

import (
	"fmt"
	"os"
	"time"
)

func (s *Source) install(
	download string,
	modified time.Time,
) error {
	stamp := modified

	if stamp.IsZero() {
		stamp = time.Unix(0, 0)
	}

	if e := os.Chtimes(download, stamp, stamp); e != nil {
		return fmt.Errorf("stamp tracker download: %w", e)
	}

	if e := os.Rename(download, s.cache); e != nil {
		return fmt.Errorf("install tracker download: %w", e)
	}

	result := s.loadCache()

	if result == nil {
		return fmt.Errorf("read tracker download: %s", s.cache)
	}

	s.tracker.Store(result)

	return nil
}
