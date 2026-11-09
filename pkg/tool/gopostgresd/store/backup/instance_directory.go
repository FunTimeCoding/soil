package backup

import "path/filepath"

func (s *Store) instanceDirectory(instance string) string {
	return filepath.Join(s.directory, instance)
}
