package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"os"
	"testing"
	"time"
)

func TestAnUnreadableChangeIsAnsweredFromTheCopyLoadedBefore(t *testing.T) {
	path := fixtureDatabase(t)
	r := memory.New()
	c := served(t, path, r)
	corrupt(t, path, time.Now().Add(time.Hour))
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
	path := fixtureDatabase(t)
	r := memory.New()
	s := openReported(t, path, nil, r)
	corrupt(t, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	assert.Count(t, 3, s.List())
	assert.Count(t, 1, r.Events())
	corrupt(t, path, time.Now().Add(2*time.Hour))
	assert.Count(t, 3, s.List())
	assert.Count(t, 2, r.Events())
}

func TestAReadableChangeAfterAnUnreadableOneIsLoaded(t *testing.T) {
	path := fixtureDatabase(t)
	s := openReported(t, path, nil, memory.New())
	corrupt(t, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	replacement := fixtureDatabase(t)
	o := openReported(t, replacement, nil, memory.New())
	_, e := o.Create("Root", "Fresh", map[string]string{})
	assert.FatalOnError(t, e)
	copyFile(t, replacement, path, time.Now().Add(2*time.Hour))
	assert.Count(t, 4, s.List())
}

func TestAWriteOnAnUnreadableChangeIsAConflict(t *testing.T) {
	path := fixtureDatabase(t)
	s := openReported(t, path, nil, memory.New())
	identifier := s.Search("forge")[0].Identifier
	corrupt(t, path, time.Now().Add(time.Hour))
	e := s.Update(
		identifier,
		map[string]string{constant.PasswordKey: "changed"},
	)
	assert.True(t, conflict.Is(e))
	password, _ := s.Reveal(identifier)
	assert.String(t, "hunter2", password)
}

func TestAConflictIsAnsweredWithoutAnEvent(t *testing.T) {
	path := fixtureDatabase(t)
	r := memory.New()
	c := served(t, path, r)
	corrupt(t, path, time.Now().Add(time.Hour))
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

func TestAMissingFileIsReportedEachTimeItGoes(t *testing.T) {
	path := fixtureDatabase(t)
	kept := join.Empty(path, ".kept")
	r := memory.New()
	s := openReported(t, path, nil, r)
	assert.FatalOnError(t, os.Rename(path, kept))
	assert.Count(t, 3, s.List())
	assert.Count(t, 1, r.Events())
	copyFile(t, kept, path, time.Now().Add(time.Hour))
	assert.Count(t, 3, s.List())
	assert.FatalOnError(t, os.Remove(path))
	assert.Count(t, 3, s.List())
	assert.Count(t, 2, r.Events())
}
