package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenameMovesFunctionsFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.False(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/old_name.go")),
	)
	assert.StringContains(t, "func NewName()", o.Read("pkg/alfa/new_name.go"))
}

func TestRenameMovesTypesFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"OldType",
		"NewType",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/new_type.go")),
	)
	assert.False(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/old_type.go")),
	)
}

func TestRenameLeavesFileWithTwoIdentities(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"PairOne",
		"PairThree",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/pair.go")),
	)
	assert.False(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/pair_three.go")),
	)
}

func TestRenameLeavesFileWhoseNameIsTaken(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	r, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Taken",
		"Busy",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/taken.go")),
	)
	assert.StringContains(t, "func Unrelated()", o.Read("pkg/alfa/busy.go"))
	assert.StringContains(
		t,
		"taken.go was not renamed: busy.go already exists",
		strings.Join(service_tester.ConcernText(r), stringsConstant.Unix),
	)
}

func TestVisibilityFlipLeavesFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	r, e := o.Symbol.ChangeVisibility(
		t.Context(),
		o.Directory,
		"Flip",
		"example/pkg/alfa",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(t, "func flip()", o.Read("pkg/alfa/flip.go"))
}

func TestRenameDryRunLeavesFile(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	_, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	assert.True(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/old_name.go")),
	)
	assert.False(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/new_name.go")),
	)
}
