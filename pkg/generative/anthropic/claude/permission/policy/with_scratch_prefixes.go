package policy

import "path/filepath"

func (w *Walk) WithScratchPrefixes(prefixes ...string) *Walk {
	for _, p := range prefixes {
		directory, okay := settle(filepath.Dir(p))

		if !okay {
			continue
		}

		w.prefixes = append(
			w.prefixes,
			filepath.Join(directory, filepath.Base(p)),
		)
	}

	return w
}
