package crap

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/crap/attribution"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
)

func printMatrix(
	m *attribution.Matrix,
	i *index.Index,
	o *option.Attribute,
) {
	single := m.Single()

	if o.Single || !o.Load {
		t := table.New(constant.ColumnTest, constant.ColumnFunction)

		for _, k := range m.SingleKeys() {
			f := i.ByKey(k)
			t.Add(
				single[k],
				join.Colon(
					system.RelativePath(o.Root, f.File),
					f.QualifiedName(),
				),
			)
		}

		t.Print()
	}

	if o.Load || !o.Single {
		t := table.New(
			constant.ColumnAlone,
			constant.ColumnTotal,
			constant.ColumnTest,
		).Right(0, 1)

		for _, l := range m.Loads() {
			t.Add(
				integers.ToString(l.Alone),
				integers.ToString(l.Total),
				l.Test,
			)
		}

		t.Print()
	}

	console.Format(
		constant.MatrixSummary,
		len(m.Tests),
		len(m.Covers),
		len(single),
	)
}
