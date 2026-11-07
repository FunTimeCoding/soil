package policy

import (
	"os"
	"path/filepath"
)

func settle(path string) (string, bool) {
	if !filepath.IsAbs(path) {
		return "", false
	}

	existing := filepath.Clean(path)
	var rest []string

	for {
		if _, e := os.Lstat(existing); e == nil {
			break
		}

		parent := filepath.Dir(existing)

		if parent == existing {
			return "", false
		}

		rest = append([]string{filepath.Base(existing)}, rest...)
		existing = parent
	}

	resolved, e := filepath.EvalSymlinks(existing)

	if e != nil {
		return "", false
	}

	return filepath.Join(append([]string{resolved}, rest...)...), true
}
