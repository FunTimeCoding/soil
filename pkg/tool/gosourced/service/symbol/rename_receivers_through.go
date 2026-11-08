package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/receiver_rename"
)

func renameReceiversThrough(
	directory string,
	packagePath string,
	renames []*receiver_rename.Rename,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := load.Packages(directory, packagePath)

	if e != nil {
		return nil, e
	}

	declaring := lookup.FindPackage(all, packagePath)

	if declaring == nil {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf("no package matches %s", packagePath),
				"",
				false,
			),
		)

		return r, nil
	}

	var planned []*plannedReceiver

	for _, n := range renames {
		planned = append(planned, planReceivers(set, declaring, n, r)...)
	}

	if commit.HasUnfixed(r) {
		return r, nil
	}

	decorations := decoration.NewSet()

	for _, p := range planned {
		if f := renameReceiver(decorations, all, set, p, r); f != nil {
			return nil, f
		}
	}

	e = decoration.Restore(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	return r, nil
}
