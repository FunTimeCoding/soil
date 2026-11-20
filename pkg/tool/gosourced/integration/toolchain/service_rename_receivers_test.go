package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/receiver_rename"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameReceiversRenamesEveryMethodOfType(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]*receiver_rename.Rename{receiver_rename.New("Shelf", "s")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.New(t).Compiles(o.Directory)
	shelf := o.Read("pkg/target/shelf.go")
	assert.StringContains(
		t,
		"func (s *Shelf) Name() string {\n\treturn s.product.Name\n}",
		shelf,
	)
	assert.StringContains(
		t,
		"func (s *Shelf) Count() int {\n\treturn s.s\n}",
		shelf,
	)
	assert.StringContains(
		t,
		"func (s Shelf) Copy() Shelf {\n\treturn s\n}",
		shelf,
	)
	assert.StringContains(
		t,
		"func (s *Shelf) Depth() int {\n\treturn s.product.y\n}",
		shelf,
	)
	assert.StringContains(t, "func (*Shelf) Kind() string {", shelf)
	assert.StringContains(t, "func (y *Other) Name() string {", shelf)
	assert.StringContains(
		t,
		"func Plain(y int) int {\n\treturn y + 1\n}",
		shelf,
	)
	test := o.Read("pkg/target/shelf_test.go")
	assert.StringContains(
		t,
		"func (s *Shelf) label() string {\n\treturn s.product.Name\n}",
		test,
	)
}

func TestRenameReceiversRefusesLocalOfName(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/local",
		[]*receiver_rename.Rename{
			receiver_rename.New("Clean", "s"),
			receiver_rename.New("Busy", "s"),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Busy.Value: s is taken in the method")
	local := o.Read("pkg/local/local.go")
	assert.StringContains(t, "func (y *Clean) Value() int {", local)
}

func TestRenameReceiversRefusesParameterOfName(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/parameter",
		[]*receiver_rename.Rename{receiver_rename.New("Busy", "s")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Busy.Scale: s is taken in the method")
}

func TestRenameReceiversRefusesOuterNameReceiverWouldShadow(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/outer",
		[]*receiver_rename.Rename{receiver_rename.New("Busy", "s")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Busy.Plus: s is taken in the method")
	outer := o.Read("pkg/outer/outer.go")
	assert.StringContains(t, "return y.value + s.value", outer)
}

func TestRenameReceiversLeavesReceiversAlreadyNamed(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]*receiver_rename.Rename{receiver_rename.New("Other", "y")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Count(t, 0, r.Entries)
}

func TestRenameReceiversRefusesBuiltInName(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]*receiver_rename.Rename{receiver_rename.New("Shelf", "len")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "len is a built-in name")
}

func TestRenameReceiversRefusesTypeWithoutMethods(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]*receiver_rename.Rename{receiver_rename.New("Product", "s")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "Product has no methods")
}

func TestRenameReceiversDryRunWritesNothing(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-receivers/src"),
	)
	r, e := o.Symbol.RenameReceivers(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		[]*receiver_rename.Rename{receiver_rename.New("Shelf", "s")},
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Count(t, 5, r.Entries)
	shelf := o.Read("pkg/target/shelf.go")
	assert.StringContains(t, "func (y *Shelf) Name() string {", shelf)
}
