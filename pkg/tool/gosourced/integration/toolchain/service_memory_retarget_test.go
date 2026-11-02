package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRenameSendsItsMovesToMemories(t *testing.T) {
	d := renameFollowFixture(t)
	s, m := service_tester.WithMemory(t)
	r, e := s.Rename(
		t.Context(),
		d,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Count(t, 1, m.Retargets)
	moves := m.Retargets[0]
	assert.Count(t, 1, moves.Paths)
	assert.String(
		t,
		filepath.Join(d, "pkg/alfa/new_name.go"),
		moves.Paths[0].To,
	)
	assert.Count(t, 1, moves.Symbols)
	assert.String(t, "NewName", moves.Symbols[0].NewName)
}

func TestRenameCommitsWhileMemoriesAreDown(t *testing.T) {
	d := renameFollowFixture(t)
	s, m := service_tester.WithMemory(t)
	m.Failure = unavailable.New("gomemoryd", time.Now(), "connection refused")
	r, e := s.Rename(
		t.Context(),
		d,
		"example/pkg/alfa",
		"OldName",
		"NewName",
		"",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.True(t, system.FileExists(filepath.Join(d, "pkg/alfa/new_name.go")))
	assert.StringContains(
		t,
		"memories not retargeted",
		strings.Join(service_tester.ConcernText(r), "\n"),
	)
}
