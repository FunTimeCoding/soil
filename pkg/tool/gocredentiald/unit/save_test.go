package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFailedSaveLeavesEntriesReadable(t *testing.T) {
	path := fixtureDatabase(t)
	c := keepass.MustNew(path, "secret")
	readOnly(t, filepath.Dir(path))
	assert.Error(t, c.Save())
	assert.String(t, "hunter2", c.ByTitle("Forge").GetPassword())
}

func TestSaveOverChangeOnDiskIsConflict(t *testing.T) {
	path := fixtureDatabase(t)
	c := keepass.MustNew(path, "secret")
	at := time.Now().Add(time.Hour)
	assert.FatalOnError(t, os.Chtimes(path, at, at))
	assert.True(t, conflict.Is(c.Save()))
}

func TestFailedWriteIsNotKeptInMemory(t *testing.T) {
	path := fixtureDatabase(t)
	s := openReported(t, path, nil, memory.New())
	identifier := s.Search("forge")[0].Identifier
	readOnly(t, filepath.Dir(path))
	assert.Error(
		t,
		s.Update(identifier, map[string]string{constant.PasswordKey: "changed"}),
	)
	password, _ := s.Reveal(identifier)
	assert.String(t, "hunter2", password)
}

func TestFailedBackupRefusesWrite(t *testing.T) {
	path := fixtureDatabase(t)
	at := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	s := openAt(t, path, at)
	identifier := s.Search("forge")[0].Identifier
	assert.FatalOnError(
		t,
		os.Mkdir(join.Empty(path, ".2026-01-02-030405.bak"), 0o700),
	)
	assert.Error(
		t,
		s.Update(identifier, map[string]string{constant.PasswordKey: "changed"}),
	)
	password, _ := openAt(t, path, at).Reveal(identifier)
	assert.String(t, "hunter2", password)
}

func TestFailedSaveLeavesNoTemporaryFile(t *testing.T) {
	path := fixtureDatabase(t)
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
