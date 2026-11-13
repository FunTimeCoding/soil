package gocheck

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
)

func openIndex(
	o *option.Check,
	graph *module_graph.Graph,
) *index.Workspace {
	if o.Workspace != nil {
		return o.Workspace(o.Root)
	}

	if graph == nil {
		graph = module_graph.BuildForest(o.Trees, o.Root)
	}

	return index.NewWithGraph(
		o.IndexDirectory(),
		o.Root,
		graph,
		goanalyze.Kinds()...,
	)
}
