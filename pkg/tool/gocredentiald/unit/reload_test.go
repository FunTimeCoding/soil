package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/database_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/unit/service_tester"
	"os"
	"testing"
	"time"
)

func TestUnreadableChangeIsAnsweredFromCopyLoadedBefore(t *testing.T) {
	s := base.New(t, database_tester.Write(t))
	c := s.Tools()
	database_tester.Corrupt(t, s.Path, time.Now().Add(time.Hour))
	assert.StringContains(
		t,
		"Forge",
		c.MustCallTool(constant.ListEntries, map[string]any{}),
	)
	assert.StringContains(
		t,
		"Forge",
		c.MustCallTool(
			constant.SearchEntries,
			map[string]any{"query": "forge"},
		),
	)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestEachUnreadableChangeIsReportedOnce(t *testing.T) {
	path := database_tester.Write(t)
	o := service_tester.New(t, path)
	s := o.Service()
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	assert.Count(t, 3, s.List())
	assert.Count(t, 1, o.Reporter.Events())
	database_tester.Corrupt(t, path, time.Now().Add(2*time.Hour))
	assert.Count(t, 3, s.List())
	assert.Count(t, 2, o.Reporter.Events())
}

func TestReadableChangeAfterUnreadableOneIsLoaded(t *testing.T) {
	path := database_tester.Write(t)
	s := service_tester.New(t, path).Service()
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	replacement := database_tester.Write(t)
	_, e := service_tester.New(t, replacement).Service().Create(
		"Root",
		"Fresh",
		map[string]string{},
	)
	assert.FatalOnError(t, e)
	database_tester.Copy(t, replacement, path, time.Now().Add(2*time.Hour))
	assert.Count(t, 4, s.List())
}

func TestWriteOnUnreadableChangeIsConflict(t *testing.T) {
	path := database_tester.Write(t)
	s := service_tester.New(t, path).Service()
	identifier := s.Search("forge")[0].Identifier
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	e := s.Update(
		identifier,
		map[string]string{constant.PasswordKey: "changed"},
	)
	assert.True(t, conflict.Is(e))
	password, _ := s.Reveal(identifier)
	assert.String(t, "hunter2", password)
}

func TestConflictIsAnsweredWithoutEvent(t *testing.T) {
	s := base.New(t, database_tester.Write(t))
	c := s.Tools()
	database_tester.Corrupt(t, s.Path, time.Now().Add(time.Hour))
	assert.StringContains(
		t,
		"refusing to write",
		c.MustCallToolError(
			constant.DeleteEntry,
			map[string]any{"identifier": "0"},
		),
	)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestMissingFileIsReportedEachTimeItGoes(t *testing.T) {
	path := database_tester.Write(t)
	kept := join.Empty(path, ".kept")
	o := service_tester.New(t, path)
	s := o.Service()
	assert.FatalOnError(t, os.Rename(path, kept))
	assert.Count(t, 3, s.List())
	assert.Count(t, 1, o.Reporter.Events())
	database_tester.Copy(t, kept, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	assert.FatalOnError(t, os.Remove(path))
	assert.Count(t, 3, s.List())
	assert.Count(t, 2, o.Reporter.Events())
}
