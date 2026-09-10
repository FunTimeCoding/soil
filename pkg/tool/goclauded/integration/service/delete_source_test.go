package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteReceiptReportsRemovedSources(t *testing.T) {
	home := t.TempDir()
	t.Setenv(constant.HomeEnvironment, home)
	identifier := "11111111-2222-3333-4444-555555555555"
	project := filepath.Join(home, ".claude", "projects", "a-project")
	errors.PanicOnError(os.MkdirAll(project, 0o755))
	transcript := filepath.Join(
		project,
		join.Empty(identifier, constant.NotationLogExtension),
	)
	errors.PanicOnError(os.WriteFile(transcript, []byte("{}\n"), 0o644))
	companion := filepath.Join(project, identifier)
	errors.PanicOnError(os.MkdirAll(companion, 0o755))
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(companion, "note.json"),
			[]byte("{}\n"),
			0o644,
		),
	)
	s := service_tester.New(t)
	s.Store.EnsureSession(identifier)
	result, e := s.Service.DeleteSession(identifier, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result.Sources)
	assert.StringContains(t, transcript, result.Sources[0])
	assert.StringContains(t, companion, result.Sources[1])
	_, f := os.Stat(transcript)
	assert.True(t, os.IsNotExist(f))
	_, g := os.Stat(companion)
	assert.True(t, os.IsNotExist(g))
}

func TestDeleteReportsNoSourcesWhenNoneExist(t *testing.T) {
	home := t.TempDir()
	t.Setenv(constant.HomeEnvironment, home)
	errors.PanicOnError(
		os.MkdirAll(filepath.Join(home, ".claude", "projects"), 0o755),
	)
	s := service_tester.New(t)
	s.Store.EnsureSession("absent")
	result, e := s.Service.DeleteSession("absent", "")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result.Sources)
}
