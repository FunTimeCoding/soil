package service

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

// A mismatch in either direction rewrites visibility silently -
// a constrained symbol lands unconstrained, or the reverse - so
// both refuse; hand-moving stays the escape hatch.
func planConstraints(
	set *token.FileSet,
	target *packages.Package,
	entries []*moveEntry,
	moveDirectory string,
) (map[string][]string, string) {
	result := make(map[string][]string)

	for _, entry := range entries {
		lines := fileConstraintLines(entry.file)

		if existing, seen := result[entry.targetFile]; seen {
			if strings.Join(existing, "\n") != strings.Join(lines, "\n") {
				return nil, fmt.Sprintf(
					"sources moving to %s carry different build constraints (%q vs %q)",
					entry.targetFile,
					strings.Join(existing, " "),
					strings.Join(lines, " "),
				)
			}

			continue
		}

		result[entry.targetFile] = lines
	}

	if target == nil {
		return result, ""
	}

	for name, lines := range result {
		file := findSyntaxFile(
			set,
			target,
			filepath.Join(moveDirectory, name),
		)

		if file == nil {
			continue
		}

		targetLines := fileConstraintLines(file)

		if strings.Join(lines, "\n") != strings.Join(targetLines, "\n") {
			return nil, fmt.Sprintf(
				"build constraint mismatch on %s: source %q, target %q - align them or move by hand",
				name,
				strings.Join(lines, " "),
				strings.Join(targetLines, " "),
			)
		}
	}

	return result, ""
}
