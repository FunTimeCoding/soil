package index

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"maps"
	"slices"
)

func (w *Workspace) ClosureKey(
	path string,
	k *kind.Kind,
) (string, bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.ensure()
	values, registered := w.facts[k.Name]

	if !registered {
		panic(
			fmt.Sprintf("kind %s is not registered with this workspace", k.Name),
		)
	}

	u, found := w.graph.Units[path]

	if !found {
		return "", false
	}

	closure := make(map[string]bool)
	queue := slices.Clone(u.Imports)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		n, workspace := w.graph.Nodes[p]

		if !workspace || closure[p] {
			continue
		}

		closure[p] = true
		queue = append(queue, n.Imports...)
	}

	h := sha256.New()
	h.Write([]byte(constant.IndexClosureKind))
	h.Write([]byte{0})
	h.Write([]byte(k.Name))

	for _, p := range slices.Sorted(maps.Keys(closure)) {
		v, known := values[p]

		if !known {
			return "", false
		}

		h.Write([]byte{0})
		h.Write([]byte(p))
		h.Write([]byte{0})
		h.Write(w.factDigest(k.Name, p, v))
	}

	return hex.EncodeToString(h.Sum(nil)), true
}
