package index

import (
	"crypto/sha256"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (w *Workspace) factDigest(
	kind string,
	path string,
	v any,
) []byte {
	cached, found := w.digests[kind]

	if !found {
		cached = make(map[string][]byte)
		w.digests[kind] = cached
	}

	if d, known := cached[path]; known {
		return d
	}

	sum := sha256.Sum256(notation.Marshal(v))
	cached[path] = sum[:]

	return sum[:]
}
