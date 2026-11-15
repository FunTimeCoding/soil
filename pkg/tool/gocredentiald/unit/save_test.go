package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/database_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFailedSaveLeavesEntriesReadable(t *testing.T) {
	path := database_tester.Write(t)
	c := keepass.MustNew(path, "secret")
	system_tester.ReadOnly(t, filepath.Dir(path))
	assert.Error(t, c.Save())
	assert.String(t, "hunter2", c.ByTitle("Forge").GetPassword())
}

func TestSaveOverChangeOnDiskIsConflict(t *testing.T) {
	path := database_tester.Write(t)
	c := keepass.MustNew(path, "secret")
	at := time.Now().Add(time.Hour)
	assert.FatalOnError(t, os.Chtimes(path, at, at))
	assert.True(t, conflict.Is(c.Save()))
}

func TestFailedWriteIsNotKeptInMemory(t *testing.T) {
	path := database_tester.Write(t)
	s := service_tester.New(t, path)
	identifier := s.Search("forge")[0].Identifier
	system_tester.ReadOnly(t, filepath.Dir(path))
	assert.Error(
		t,
		s.Update(
			identifier,
			map[string]string{constant.PasswordKey: "changed"},
		),
	)
	password, _ := s.Reveal(identifier)
	assert.String(t, "hunter2", password)
}

func TestFailedBackupRefusesWrite(t *testing.T) {
	path := database_tester.Write(t)
	at := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	s := service_tester.New(t, path, service_tester.WithClock(at))
	identifier := s.Search("forge")[0].Identifier
	assert.FatalOnError(
		t,
		os.Mkdir(join.Empty(path, ".2026-01-02-030405.bak"), 0o700),
	)
	assert.Error(
		t,
		s.Update(
			identifier,
			map[string]string{constant.PasswordKey: "changed"},
		),
	)
	password, _ := service_tester.New(
		t,
		path,
		service_tester.WithClock(at),
	).Reveal(
		identifier,
	)
	assert.String(t, "hunter2", password)
}

func TestFailedSaveLeavesNoTemporaryFile(t *testing.T) {
	path := database_tester.Write(t)
	c := keepass.MustNew(path, "secret")
	i, e := os.Stat(path)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, os.Remove(path))
	assert.FatalOnError(t, os.MkdirAll(filepath.Join(path, "occupied"), 0o700))
	assert.FatalOnError(t, os.Chtimes(path, i.ModTime(), i.ModTime()))
	assert.Error(t, c.Save())
	_, e = os.Stat(join.Empty(path, ".tmp"))
	assert.True(t, os.IsNotExist(e))
}
