package index

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"runtime"
	"slices"
)

func unitDigest(
	g *module_graph.Graph,
	u *module_graph.Node,
	files string,
	fingerprints map[string]string,
	kind string,
) string {
	h := sha256.New()
	write := func(s ...string) {
		for _, v := range s {
			h.Write([]byte(v))
			h.Write([]byte{0})
		}
	}
	write(formatVersion(), kind, runtime.Version(), u.Path, files)
	write(g.Tags...)

	for _, i := range slices.Compact(slices.Sorted(slices.Values(u.Imports))) {
		if _, workspace := g.Nodes[i]; workspace {
			write(i, fingerprints[i])

			continue
		}

		write(i, externalVersion(i, g.Requirements))
	}

	return hex.EncodeToString(h.Sum(nil))
}
