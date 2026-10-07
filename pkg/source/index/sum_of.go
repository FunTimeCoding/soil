package index

import (
	"crypto/sha256"
	"github.com/funtimecoding/soil/pkg/source/types/file_sum"
	"os"
)

func (w *Workspace) sumOf(path string) ([32]byte, bool) {
	i, e := os.Stat(path)

	if e != nil {
		return [32]byte{}, false
	}

	if cached := w.sums[path]; cached != nil &&
		cached.Size == i.Size() &&
		cached.Modified.Equal(i.ModTime()) {
		return cached.Sum, true
	}

	content, f := os.ReadFile(path)

	if f != nil {
		return [32]byte{}, false
	}

	result := sha256.Sum256(content)
	w.sums[path] = file_sum.New(i.Size(), i.ModTime(), result)

	return result, true
}
