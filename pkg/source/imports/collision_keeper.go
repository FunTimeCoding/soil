package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func collisionKeeper(
	v *import_view.View,
	declaredName string,
	group []*import_entry.Entry,
) (string, string) {
	if declaredName == constant.MockPackage {
		return "", constant.KeeperMock
	}

	var paths []string
	plain := ""
	standard := make(map[string]bool)

	for _, e := range group {
		paths = append(paths, e.Path)
		standard[e.Path] = e.Standard

		if e.Local == declaredName {
			plain = e.Path
		}
	}

	return Keeper(
		v.PackagePath,
		declaredName,
		paths,
		plain,
		func(p string) bool { return standard[p] },
	)
}
