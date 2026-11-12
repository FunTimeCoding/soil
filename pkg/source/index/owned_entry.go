package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"slices"
)

func ownedEntry(name string) bool {
	return constant.GenerationPattern.MatchString(name) ||
		constant.BuildPattern.MatchString(name) ||
		slices.Contains(
			[]string{
				constant.IndexPackageKind,
				constant.IndexImportsKind,
				constant.IndexReferencesKind,
				constant.IndexFactsKind,
				constant.IndexExternalKind,
			},
			name,
		)
}
