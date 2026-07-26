package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"sort"
)

func restoreDecorations(
	decorations *decoration.Set,
	names *resolve.Names,
	skip map[string]bool,
) error {
	var filenames []string

	for filename := range decorations.Files {
		if !skip[filename] {
			filenames = append(filenames, filename)
		}
	}

	sort.Strings(filenames)

	for _, filename := range filenames {
		file := decorations.Files[filename]
		e := restoreDecoratedFile(
			names,
			decorations.PackagePaths[file],
			decorations.Aliases[file],
			file,
			filename,
		)

		if e != nil {
			return e
		}
	}

	return nil
}
