package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func planConstraints(
	set *token.FileSet,
	target *packages.Package,
	entries []*relocation.Entry,
	moveDirectory string,
) (map[string][]string, string) {
	result := make(map[string][]string)

	for _, entry := range entries {
		lines := fileConstraintLines(entry.File)

		if existing, seen := result[entry.TargetFile]; seen {
			if join.NewLine(existing) != join.NewLine(lines) {
				return nil, fmt.Sprintf(
					"sources moving to %s carry different build constraints (%q vs %q)",
					entry.TargetFile,
					join.Space(existing...),
					join.Space(lines...),
				)
			}

			continue
		}

		result[entry.TargetFile] = lines
	}

	if target == nil {
		return result, ""
	}

	for name, lines := range result {
		file := findSyntaxFile(set, target, filepath.Join(moveDirectory, name))

		if file == nil {
			continue
		}

		targetLines := fileConstraintLines(file)

		if join.NewLine(lines) != join.NewLine(targetLines) {
			return nil, fmt.Sprintf(
				"build constraint mismatch on %s: source %q, target %q - align them or move by hand",
				name,
				join.Space(lines...),
				join.Space(targetLines...),
			)
		}
	}

	return result, ""
}
