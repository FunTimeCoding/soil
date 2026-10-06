package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"slices"
	"strings"
)

func reachModule(
	indexDirectory string,
	directory string,
	renames map[string]violation,
) (*reachedModule, error) {
	i := index.New(indexDirectory, directory).References()

	if unindexed := i.Unindexed(); len(unindexed) > 0 {
		return nil, fmt.Errorf(
			"%s does not type-check",
			strings.Join(unindexed, ", "),
		)
	}

	names := make(map[string]bool, len(renames))
	var patterns []string

	for t, v := range renames {
		names[v.ident.Name] = true

		for _, unit := range i.Referencing(t) {
			patterns = append(patterns, i.Directory(unit))
		}
	}

	if len(patterns) == 0 {
		return nil, nil
	}

	patterns = slices.Compact(slices.Sorted(slices.Values(patterns)))
	all, set, e := resolve.LoadPackages(directory, patterns...)

	if e != nil {
		return nil, e
	}

	result := &reachedModule{directory: directory, fileSet: set}
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		for ident, use := range p.TypesInfo.Uses {
			if !names[ident.Name] || seen[ident.Pos()] {
				continue
			}

			t, okay := xref.Target(use)

			if !okay {
				continue
			}

			v, found := renames[t]

			if !found {
				continue
			}

			seen[ident.Pos()] = true
			replacement := segment.ReplaceSegment(ident.Name, v.segment, v.fix)
			result.edits = append(
				result.edits,
				edit{
					position: ident.Pos(),
					end:      ident.End(),
					newText:  replacement,
				},
			)
			position := set.Position(ident.Pos())
			result.concerns = append(
				result.concerns,
				concern.NewLine(
					"renamed",
					fmt.Sprintf("renamed %s → %s", ident.Name, replacement),
					position.Filename,
					position.Line,
					"",
					true,
				),
			)
		}
	}

	return result, nil
}
