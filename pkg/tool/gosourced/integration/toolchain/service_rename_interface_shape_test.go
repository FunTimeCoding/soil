package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameOfInterfaceMethodDropsShapeOnlySatisfaction(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-interface-shape/src"),
	)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/shape",
		"Close",
		"Shut",
		"Closer",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func (d *Door) Shut() error",
		o.Read("pkg/door/door.go"),
	)
	var dropped []string

	for _, c := range r.Entries {
		if c.Key == constant.ConcernDroppedSatisfaction {
			dropped = append(dropped, c.Text)
		}
	}

	assert.Any(
		t,
		[]string{
			"example/pkg/door.Door satisfies io.Closer outside the workspace by shape only - no package holding it reaches io, so the rename drops that satisfaction",
		},
		dropped,
	)
}

func TestRenameOfInterfaceMethodBlockedWhereHolderReachesInterface(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-interface-shape/src"),
	)
	system_tester.Write(
		t,
		o.Directory,
		"pkg/use/drain.go",
		"package use\n\nimport \"example/pkg/stream\"\n\nvar _ = stream.Drain\n",
	)
	before := rewrite_tester.DirectoryTree(t, o.Directory)
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/shape",
		"Close",
		"Shut",
		"Closer",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(
		t,
		r,
		"example/pkg/door.Door satisfies io.Closer outside the workspace, and example/pkg/use reaches io",
	)
	assert.String(t, before, rewrite_tester.DirectoryTree(t, o.Directory))
}
