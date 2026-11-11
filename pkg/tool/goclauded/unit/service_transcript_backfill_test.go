package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestBackfillSkipsSessionWithoutTranscript(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("present-session")
	s.Store.EnsureSession("absent-session")
	s.WriteContextLoadFile("present-session")
	result := s.Transcript.BackfillAllSessions()
	assert.Integer(t, 1, result.Enriched)
	assert.Integer(t, 1, result.Skipped)
}

func TestColdBackfillSkipsSessionWithoutTranscript(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("present-session")
	s.Store.EnsureSession("absent-session")
	s.WriteContextLoadFile("present-session")
	result := s.Transcript.ColdBackfillAllSessions()
	assert.Integer(t, 1, result.Enriched)
	assert.Integer(t, 1, result.Skipped)
}
