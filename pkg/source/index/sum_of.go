package index

import (
	"crypto/sha256"
	"os"
)

func (w *Workspace) sumOf(path string) ([32]byte, bool) {
	i, e := os.Stat(path)

	if e != nil {
		return [32]byte{}, false
	}

	if cached := w.sums[path]; cached != nil &&
		cached.size == i.Size() &&
		cached.modified.Equal(i.ModTime()) {
		return cached.sum, true
	}

	content, f := os.ReadFile(path)

	if f != nil {
		return [32]byte{}, false
	}

	result := sha256.Sum256(content)
	w.sums[path] = &fileSum{size: i.Size(), modified: i.ModTime(), sum: result}

	return result, true
}
