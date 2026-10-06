package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"runtime"
	"slices"
)

func key(
	g *module_graph.Graph,
	n *module_graph.Node,
	files string,
	fingerprints map[string]string,
) string {
	h := sha256.New()
	write := func(s ...string) {
		for _, v := range s {
			h.Write([]byte(v))
			h.Write([]byte{0})
		}
	}
	write(formatVersion(), runtime.Version(), n.Path, files)
	write(g.Tags...)

	for _, i := range slices.Sorted(slices.Values(n.Imports)) {
		if _, workspace := g.Nodes[i]; workspace {
			write(i, fingerprints[i])

			continue
		}

		write(i, externalVersion(i, g.Requirements))
	}

	return hex.EncodeToString(h.Sum(nil))
}
