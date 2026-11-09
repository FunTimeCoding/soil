package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestBackupIsWrittenWithItsChecksum(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	s.Dump.Tables = 7
	b, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "2030-01-02T03-04-05", b.Name)
	assert.Integer(t, 7, b.Tables)
	assert.String(
		t,
		"a405eba78bf2e6db44ebe0b28bbc9cdc449f9ac990d2029c50a15e6853cfdf20",
		b.Checksum,
	)
	line, f := os.ReadFile(
		filepath.Join(s.Directory, "alfa", "2030-01-02T03-04-05.dump.sha256"),
	)
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"a405eba78bf2e6db44ebe0b28bbc9cdc449f9ac990d2029c50a15e6853cfdf20  2030-01-02T03-04-05.dump\n",
		string(line),
	)
}

func TestFailedDumpLeavesNoFile(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	s.Dump.Failure = unexpected.Format("pg_dump: connection refused")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.True(t, e != nil)
	listed, f := s.Backup.List("alfa")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, listed)
}

func TestBackupsListNewestFirst(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(s.Directory, "alfa", "2029-12-31T23-59-59.dump"),
			[]byte("older"),
			0o640,
		),
	)
	listed, f := s.Backup.List("alfa")
	assert.FatalOnError(t, f)
	assert.Count(t, 2, listed)
	assert.String(t, "2030-01-02T03-04-05", listed[0].Name)
	assert.String(t, "2029-12-31T23-59-59", listed[1].Name)
}

func TestDeleteRemovesBackupAndItsChecksum(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.Backup.Delete("alfa", "2030-01-02T03-04-05"))
	entries, f := os.ReadDir(filepath.Join(s.Directory, "alfa"))
	assert.FatalOnError(t, f)
	assert.Count(t, 0, entries)
}

func TestBackupNameThatIsNotTimestampIsRefused(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	assert.True(t, validation.Is(s.Backup.Delete("alfa", "../alfa")))
}

func TestAbsentBackupIsNotFound(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	_, e := s.Backup.Restore(t.Context(), "alfa", "2030-01-02T03-04-05", "")
	assert.True(t, not_found.Is(e))
}

func TestUnknownInstanceIsNotFound(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "zulu")
	assert.True(t, not_found.Is(e))
}

func TestRestoreDefaultsToNewDatabaseBeside(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	target, f := s.Backup.Restore(
		t.Context(),
		"alfa",
		"2030-01-02T03-04-05",
		"",
	)
	assert.FatalOnError(t, f)
	assert.String(t, "alfa_restore_2030-01-02T03-04-05", target)
	statements := clients["alfa"].Statements()
	assert.String(
		t,
		`CREATE DATABASE "alfa_restore_2030-01-02T03-04-05"`,
		statements[len(statements)-1].SQL,
	)
	assert.Count(t, 1, s.Dump.Restored)
	assert.String(t, target, s.Dump.Restored[0].Database)
	assert.False(t, s.Dump.Restored[0].Clean)
}

func TestRestoreRefusesExistingOtherDatabase(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	clients["alfa"].SeedResult(
		constant.DatabaseExistsQuery,
		result.New([]string{"?column?"}, [][]any{{1}}),
	)
	_, f := s.Backup.Restore(
		t.Context(),
		"alfa",
		"2030-01-02T03-04-05",
		"bravo",
	)
	assert.True(t, conflict.Is(f))
	assert.Count(t, 0, s.Dump.Restored)
}

func TestRestoreInPlaceIsRefusedWhileOthersAreConnected(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	clients["alfa"].SeedResult(
		constant.ConnectedSessionsQuery,
		result.New(
			nil,
			[][]any{{"", "alfa", "10.0.0.5", "2030-01-02T00:00:00Z"}},
		),
	)
	_, f := s.Backup.Restore(t.Context(), "alfa", "2030-01-02T03-04-05", "alfa")
	assert.True(t, conflict.Is(f))
	assert.StringContains(
		t,
		"(no application name) (alfa@10.0.0.5 since 2030-01-02T00:00:00Z)",
		f.Error(),
	)
	assert.Count(t, 0, s.Dump.Restored)
}

func TestRestoreInPlaceLeavesItsOwnSessionsOut(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Backup.Restore(t.Context(), "alfa", "2030-01-02T03-04-05", "alfa")
	assert.FatalOnError(t, f)
	statements := clients["alfa"].Statements()
	assert.StringContains(t, "FROM pg_stat_activity", statements[0].SQL)
	assert.Any(t, []any{"gopostgresd%"}, statements[0].Arguments)
}

func TestRestoreInPlaceCleansAndRebuildsPool(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	_, e := s.Backup.Create(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Backup.Restore(t.Context(), "alfa", "2030-01-02T03-04-05", "alfa")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, s.Dump.Restored)
	assert.True(t, s.Dump.Restored[0].Clean)
	assert.Integer(t, 1, clients["alfa"].Closed)
}
