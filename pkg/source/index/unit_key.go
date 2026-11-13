package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (w *Workspace) UnitKey(path string) (string, bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.ensure()
	u, found := w.graph.Units[path]

	if !found {
		return "", false
	}

	for _, i := range u.Imports {
		if _, workspace := w.graph.Nodes[i]; !workspace {
			continue
		}

		if _, known := w.fingerprints[i]; !known {
			return "", false
		}
	}

	sum := sha256.Sum256(
		[]byte(
			join.Space(
				unitDigest(
					w.graph,
					u,
					w.fileHash(u.Files),
					w.fingerprints,
					constant.IndexUnitKind,
				),
				w.build,
			),
		),
	)

	return hex.EncodeToString(sum[:]), true
}
