package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/transcript_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestOversizedLineDoesNotStopExtraction(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("oversized-session")
	body := join.Empty(
		fixture.Read("claude", "context-loads.jsonl"),
		transcript_tester.OversizedLine(),
		"\n",
	)
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(
				s.Harbor,
				join.Empty("oversized-session", soil.NotationLogExtension),
			),
			[]byte(body),
			0o644,
		),
	)
	s.Transcript.EnrichSession("oversized-session")
	loads, e := s.Transcript.SessionContextLoads(
		t.Context(),
		"oversized-session",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 11, len(loads))
}
