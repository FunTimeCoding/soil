package gocheck

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/clean_cache"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
)

func typedStages(
	o *option.Check,
	patterns []string,
	parent *layer.Layer,
	graph *module_graph.Graph,
) (*layer.Layer, []*output.Results, []string) {
	handle := openIndex(o, graph)
	graph = handle.Graph()
	cache := clean_cache.New(
		handle,
		parent,
		configuration.Load(o.Root).Comment,
		o.ModuleInventory().ModulePaths(),
	)
	changed := changedDirectories(o.Root, parent.Changes())
	units := patternUnits(graph, o.Root, patterns)
	keys := make(map[string]string)
	dirty := make(map[string]bool)
	all := make(map[string]bool)
	var checked []*module_graph.Node

	for _, u := range units {
		relative := relativeDirectory(o.Root, u.Directory)
		all[relative] = true
		k, keyed := cache.Key(u)

		if keyed {
			keys[u.Path] = k
		}

		if keyed && !o.Fresh && !changed[relative] && cache.Holds(k) {
			continue
		}

		dirty[relative] = true
		checked = append(checked, u)
	}

	if len(units) > 0 && len(dirty) == 0 {
		return parent, cleanSections(o.Root, patterns), nil
	}

	work := patterns

	if 2*len(dirty) <= len(all) {
		work = directoryPatterns(dirty)
	}

	loaded, set, e := resolve.LoadPackagesTagged(
		o.Root,
		graph.Tags,
		parent.Overlay(),
		work...,
	)
	errors.PanicOnError(e)

	if failures := loadFailures(o.Root, loaded); len(failures.Entries) > 0 {
		failures.Header = lint.HeaderLine(
			constant.Identity.Name(),
			o.Root,
			lint.PatternDetail(patterns),
		)

		return parent, []*output.Results{failures}, directoryList(dirty)
	}

	reported := resolve.PackagePaths(loaded)
	w, fixed := gofixStage(
		o,
		work,
		patterns,
		reported,
		handle,
		loaded,
		set,
		parent,
	)

	if w == nil {
		return parent, []*output.Results{fixed}, directoryList(dirty)
	}

	sections := []*output.Results{
		fixed,
		goanalyzeStage(o, work, patterns, reported, handle, w),
	}
	markClean(cache, o.Root, checked, keys, sections, w.Layer().Changes())

	return w.Layer(), sections, directoryList(dirty)
}
