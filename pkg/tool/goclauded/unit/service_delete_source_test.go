package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteReceiptReportsRemovedSources(t *testing.T) {
	s := service_tester.New(t)
	identifier := "11111111-2222-3333-4444-555555555555"
	transcript := filepath.Join(
		s.Source,
		join.Empty(identifier, soil.NotationLogExtension),
	)
	errors.PanicOnError(os.WriteFile(transcript, []byte("{}\n"), 0o644))
	companion := filepath.Join(s.Source, identifier)
	errors.PanicOnError(os.MkdirAll(companion, 0o755))
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(companion, "note.json"),
			[]byte("{}\n"),
			0o644,
		),
	)
	s.Store.EnsureSession(identifier)
	result, e := s.Session.DeleteSession(identifier, "")
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
	s := service_tester.New(t)
	s.Store.EnsureSession("absent")
	result, e := s.Session.DeleteSession("absent", "")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result.Sources)
}
