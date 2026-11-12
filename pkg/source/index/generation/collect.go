package generation

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"path/filepath"
	"strings"
	"time"
)

func Collect(
	root string,
	keep string,
	now time.Time,
	retention time.Duration,
	owns func(string) bool,
) []string {
	var result []string

	for _, sibling := range siblings(root, keep) {
		if strings.HasPrefix(sibling, constant.GenerationTrash) {
			discard(root, sibling)

			continue
		}

		if !owns(sibling) {
			continue
		}

		if now.Sub(touched(filepath.Join(root, sibling), now)) >= retention {
			discard(root, sibling)
			result = append(result, sibling)
		}
	}

	return result
}
