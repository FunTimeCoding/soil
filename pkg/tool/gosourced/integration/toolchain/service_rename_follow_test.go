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

func TestRenameMovesAFunctionsFile(t *testing.T) {
	d := renameFollowFixture(t)
	r, e := service_tester.NewService().Rename(
		d,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.False(t, system.FileExists(filepath.Join(d, "pkg/alfa/old_name.go")))
	assert.StringContains(
		t,
		"func NewName()",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/new_name.go"),
	)
}

func TestRenameMovesATypesFile(t *testing.T) {
	d := renameFollowFixture(t)
	r, e := service_tester.NewService().Rename(
		d,
		"example/pkg/alfa",
		"OldType",
		"NewType",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, system.FileExists(filepath.Join(d, "pkg/alfa/new_type.go")))
	assert.False(t, system.FileExists(filepath.Join(d, "pkg/alfa/old_type.go")))
}

func TestRenameLeavesAFileWithTwoIdentities(t *testing.T) {
	d := renameFollowFixture(t)
	r, e := service_tester.NewService().Rename(
		d,
		"example/pkg/alfa",
		"PairOne",
		"PairThree",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, system.FileExists(filepath.Join(d, "pkg/alfa/pair.go")))
	assert.False(
		t,
		system.FileExists(filepath.Join(d, "pkg/alfa/pair_three.go")),
	)
}

func TestRenameLeavesAFileWhoseNameIsTaken(t *testing.T) {
	d := renameFollowFixture(t)
	r, e := service_tester.NewService().Rename(
		d,
		"example/pkg/alfa",
		"Taken",
		"Busy",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, system.FileExists(filepath.Join(d, "pkg/alfa/taken.go")))
	assert.StringContains(
		t,
		"func Unrelated()",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/busy.go"),
	)
	assert.StringContains(
		t,
		"taken.go was not renamed: busy.go already exists",
		strings.Join(service_tester.ConcernText(r), stringsConstant.Unix),
	)
}

func TestAVisibilityFlipLeavesTheFile(t *testing.T) {
	d := renameFollowFixture(t)
	r, e := service_tester.NewService().ChangeVisibility(
		d,
		"Flip",
		"example/pkg/alfa",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"func flip()",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/flip.go"),
	)
}

func TestARenameDryRunLeavesTheFile(t *testing.T) {
	d := renameFollowFixture(t)
	_, e := service_tester.NewService().Rename(
		d,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	assert.True(t, system.FileExists(filepath.Join(d, "pkg/alfa/old_name.go")))
	assert.False(t, system.FileExists(filepath.Join(d, "pkg/alfa/new_name.go")))
}
