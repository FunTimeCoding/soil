package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRenameSendsItsMovesToMemories(t *testing.T) {
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
	assert.Count(t, 1, o.Memory.Retargets)
	moves := o.Memory.Retargets[0]
	assert.Count(t, 1, moves.Paths)
	assert.String(
		t,
		filepath.Join(o.Directory, "pkg/alfa/new_name.go"),
		moves.Paths[0].New,
	)
	assert.Count(t, 1, moves.Symbols)
	assert.String(t, "NewName", moves.Symbols[0].NewName)
}

func TestMemoryFailurePeerCapturedIsNotCapturedAgain(t *testing.T) {
	p := base.New(t)
	o := service_tester.New(
		t,
		service_tester.WithFixture("rename-follow/src"),
		service_tester.WithMemoryPeer(p.MemoryClient()),
	)
	result, e := o.Symbol.Rename(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, p.Reporter.Events())
	assert.Count(t, 0, o.Reporter.Events())
	assert.StringContains(
		t,
		"memories not retargeted",
		strings.Join(service_tester.ConcernText(result), "\n"),
	)
}

func TestRenameCommitsWhileMemoriesAreDown(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("rename-follow/src"))
	o.Memory.Failure = unavailable.New(
		"gomemoryd",
		time.Now(),
		"connection refused",
	)
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
	assert.True(
		t,
		system.FileExists(filepath.Join(o.Directory, "pkg/alfa/new_name.go")),
	)
	assert.StringContains(
		t,
		"memories not retargeted",
		strings.Join(service_tester.ConcernText(r), "\n"),
	)
}
