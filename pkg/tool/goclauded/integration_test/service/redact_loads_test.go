package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestRedactedLoadsHideNameAndReference(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("redacted-session")
	s.Memory.Redacted = map[int64]bool{109: true}
	writeContextLoadFile(t, s.Harbor, "redacted-session")
	s.Service.EnrichSession("redacted-session")
	loads, e := s.Service.ContextLoadsBySession("redacted-session")
	assert.FatalOnError(t, e)
	seen := false

	for _, entry := range loads {
		if entry.Name == "kilo" {
			seen = true
		}

		assert.True(t, entry.Reference != "109")
	}

	assert.True(t, !seen)
}

func TestUnredactedLoadsKeepTheirNames(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("plain-session")
	writeContextLoadFile(t, s.Harbor, "plain-session")
	s.Service.EnrichSession("plain-session")
	loads, e := s.Service.ContextLoadsBySession("plain-session")
	assert.FatalOnError(t, e)
	seen := false

	for _, entry := range loads {
		if entry.Name == "kilo" {
			seen = true
		}
	}

	assert.True(t, seen)
}
