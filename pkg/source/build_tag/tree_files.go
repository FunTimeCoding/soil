package build_tag

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"path/filepath"
)

func TreeFiles(t *warm_tree.Tree) map[string][]string {
	result := make(map[string][]string)

	for _, p := range t.Files() {
		absolute := filepath.Join(t.Root(), filepath.FromSlash(p))

		if !Scanned(t.Root(), absolute) {
			continue
		}

		b, e := t.Read(p)

		if e != nil {
			continue
		}

		if tags := extractFrom(bytes.NewReader(b)); len(tags) > 0 {
			result[absolute] = tags
		}
	}

	return result
}
