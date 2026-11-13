package gocheck

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/clean_cache"
)

func markClean(
	cache *clean_cache.Cache,
	root string,
	checked []*module_graph.Node,
	keys map[string]string,
	sections []*output.Results,
	changes []*layer_change.Change,
) {
	findings := findingDirectories(sections)
	changed := changedDirectories(root, changes)

	for _, u := range checked {
		relative := relativeDirectory(root, u.Directory)
		before, keyed := keys[u.Path]

		if !keyed || findings[relative] || changed[relative] {
			continue
		}

		if after, okay := cache.Key(u); okay && after == before {
			cache.Add(before)
		}
	}
}
