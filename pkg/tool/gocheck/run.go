package gocheck

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"slices"
)

func Run(o *option.Check) (*check_result.Result, error) {
	result := check_result.New(o.Root)
	patterns := []string{goModConstant.AllPackages}
	var scopes []string
	var graph *module_graph.Graph

	if o.Transcript != "" {
		r, announcement, e := reached(o)

		if e != nil {
			return nil, e
		}

		result.Announcement = announcement

		if len(r.Packages) == 0 && len(r.Documents) == 0 {
			result.Unchanged = true

			return result, nil
		}

		patterns = r.Packages
		scopes = slices.Concat(r.Packages, r.Documents)
		graph = r.Graph
	}

	replacing := o.ModuleInventory().Replacing(o.Root)
	fixes := layer.New(
		layer.New(nil, slices.Concat([]string{o.Root}, replacing)...),
	)
	result.Sections = append(result.Sections, golintStage(o, scopes, fixes))
	result.Sections = append(result.Sections, goauditStage(o)...)
	top := fixes

	if len(patterns) > 0 {
		var typed []*output.Results
		top, typed, result.Checked = typedStages(o, patterns, fixes, graph)
		result.Sections = append(result.Sections, typed...)
	}

	finishStack(o, top, result)

	return result, nil
}
