package module_graph

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/build"
	"maps"
	"path/filepath"
	"slices"
)

func Update(
	previous *Graph,
	root string,
	changed []string,
) *Graph {
	for _, p := range changed {
		if name := filepath.Base(p); name == constant.ModFile ||
			name == constant.SumFile {
			return Build(root)
		}
	}

	fileTags := maps.Clone(previous.FileTags)

	for _, p := range changed {
		if !build_tag.Scanned(root, p) {
			continue
		}

		if tags := build_tag.Extract(p); len(tags) > 0 {
			fileTags[p] = tags
		} else {
			delete(fileTags, p)
		}
	}

	if !slices.Equal(build_tag.Union(fileTags), previous.Tags) {
		return Build(root)
	}

	c := build.Default
	c.BuildTags = previous.Tags
	module, requirements, replaced := readModule(root)
	nodes := maps.Clone(previous.Nodes)
	units := maps.Clone(previous.Units)
	directories := make(map[string]bool)

	for _, p := range changed {
		directories[filepath.Dir(p)] = true
	}

	for directory := range directories {
		if p, inside := mainPath(root, module, directory); inside {
			delete(nodes, p)
			delete(units, p)
			delete(units, join.Empty(p, "_test"))
			n, directoryUnits := importDirectory(&c, directory, p)

			if n != nil {
				nodes[p] = n
			}

			for _, u := range directoryUnits {
				units[u.Path] = u
			}

			continue
		}

		if p, inside := replacedPath(replaced, directory); inside {
			delete(nodes, p)
		}
	}

	reach(&c, module, nodes, units, replaced)

	return NewGraph(nodes, units, fileTags, requirements)
}
