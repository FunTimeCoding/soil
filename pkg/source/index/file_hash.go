package index

import (
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
	"slices"
)

func (w *Workspace) fileHash(files []string) string {
	sorted := slices.Sorted(slices.Values(files))
	h := sha256.New()

	for _, f := range sorted {
		s, okay := w.sumOf(f)

		if !okay {
			continue
		}

		h.Write([]byte(filepath.Base(f)))
		h.Write([]byte{0})
		h.Write(s[:])
	}

	return hex.EncodeToString(h.Sum(nil))
}
