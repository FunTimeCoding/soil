package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
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
	path := database_tester.Write(t)
	r := memory.New()
	c := base.New(t, path, r).Connect(t)
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	assert.StringContains(
		t,
		"Forge",
		c.MustCallTool(constant.ListEntries, map[string]any{}),
	)
	assert.StringContains(
		t,
		"Forge",
		c.MustCallTool(constant.SearchEntries, map[string]any{"query": "forge"}),
	)
	assert.Count(t, 1, r.Events())
}

func TestEachUnreadableChangeIsReportedOnce(t *testing.T) {
	path := database_tester.Write(t)
	r := memory.New()
	s := service_tester.New(t, path, service_tester.WithReporter(r))
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	assert.Count(t, 3, s.List())
	assert.Count(t, 1, r.Events())
	database_tester.Corrupt(t, path, time.Now().Add(2*time.Hour))
	assert.Count(t, 3, s.List())
	assert.Count(t, 2, r.Events())
}

func TestReadableChangeAfterUnreadableOneIsLoaded(t *testing.T) {
	path := database_tester.Write(t)
	s := service_tester.New(t, path)
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	replacement := database_tester.Write(t)
	o := service_tester.New(t, replacement)
	_, e := o.Create("Root", "Fresh", map[string]string{})
	assert.FatalOnError(t, e)
	database_tester.Copy(t, replacement, path, time.Now().Add(2*time.Hour))
	assert.Count(t, 4, s.List())
}

func TestWriteOnUnreadableChangeIsConflict(t *testing.T) {
	path := database_tester.Write(t)
	s := service_tester.New(t, path)
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
	path := database_tester.Write(t)
	r := memory.New()
	c := base.New(t, path, r).Connect(t)
	database_tester.Corrupt(t, path, time.Now().Add(time.Hour))
	assert.StringContains(
		t,
		"refusing to write",
		c.MustCallToolError(
			constant.DeleteEntry,
			map[string]any{"identifier": "0"},
		),
	)
	assert.Count(t, 1, r.Events())
}

func TestMissingFileIsReportedEachTimeItGoes(t *testing.T) {
	path := database_tester.Write(t)
	kept := join.Empty(path, ".kept")
	r := memory.New()
	s := service_tester.New(t, path, service_tester.WithReporter(r))
	assert.FatalOnError(t, os.Rename(path, kept))
	assert.Count(t, 3, s.List())
	assert.Count(t, 1, r.Events())
	database_tester.Copy(t, kept, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	assert.FatalOnError(t, os.Remove(path))
	assert.Count(t, 3, s.List())
	assert.Count(t, 2, r.Events())
}
