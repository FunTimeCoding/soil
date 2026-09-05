package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func oversizedLine() string {
	return join.Empty(
		`{"type":"assistant","timestamp":"2026-08-16T00:30:00.000Z",`,
		`"message":{"role":"assistant","content":[{"type":"text","text":"`,
		strings.Repeat("x", 2*1024*1024),
		`"}]}}`,
	)
}

func TestOversizedLineDoesNotStopExtraction(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("oversized-session")
	body := join.Empty(
		fixture.Read("claude", "context-loads.jsonl"),
		oversizedLine(),
		"\n",
	)
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(
				s.Harbor,
				join.Empty("oversized-session", constant.NotationLogExtension),
			),
			[]byte(body),
			0o644,
		),
	)
	s.Service.EnrichSession("oversized-session")
	loads, e := s.Service.ContextLoadsBySession("oversized-session")
	assert.FatalOnError(t, e)
	assert.Integer(t, 11, len(loads))
}
