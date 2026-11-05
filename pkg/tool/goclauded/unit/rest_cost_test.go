package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"path/filepath"
	"testing"
	"time"
)

func TestCostCarriesItsWindowAndFlagsAnUnknownModel(t *testing.T) {
	s := base.New(t)
	at := time.Now().UTC().Add(-time.Minute).Format(time.RFC3339Nano)
	line := `{"uuid":"%s","type":"assistant","timestamp":"%s","requestId":"%s","message":{"id":"%s","role":"assistant","model":"%s","content":[],"usage":{"input_tokens":10,"output_tokens":5}}}`
	system.WriteFile(
		filepath.Join(s.Harbor, "delta.jsonl"),
		[]byte(
			join.Empty(
				fmt.Sprintf(line, "d1", at, "r1", "m1", "claude-sonnet-4-5"),
				"\n",
				fmt.Sprintf(line, "d2", at, "r2", "m2", "kilo-model"),
				"\n",
			),
		),
		0644,
	)
	days := 2
	r, e := s.RESTClient(t).GetCostWithResponse(
		context.Background(),
		&client.GetCostParams{Days: &days},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.Integer(t, 2, r.JSON200.Days)
	assert.FatalNotNil(t, r.JSON200.Models)
	unknown := map[string]bool{}

	for _, m := range *r.JSON200.Models {
		unknown[m.Model] = m.UnknownModel
	}

	assert.Count(t, 2, unknown)
	assert.False(t, unknown["sonnet"])
	assert.True(t, unknown["kilo-model"])
}

func TestSessionDetailFlagsAnUnknownModel(t *testing.T) {
	s := base.New(t)
	user := `{"type":"user","timestamp":"2026-05-21T10:00:00Z","sessionId":"echo","slug":"echo-slug","cwd":"/home/user","message":{"role":"user","content":"alfa"}}`
	line := `{"uuid":"%s","type":"assistant","timestamp":"2026-05-21T10:01:00Z","sessionId":"echo","requestId":"%s","message":{"id":"%s","role":"assistant","model":"%s","content":[],"usage":{"input_tokens":10,"output_tokens":5}}}`
	system.WriteFile(
		filepath.Join(s.Harbor, "echo.jsonl"),
		[]byte(
			join.Empty(
				user,
				"\n",
				fmt.Sprintf(line, "e1", "r1", "m1", "claude-sonnet-4-5"),
				"\n",
				fmt.Sprintf(line, "e2", "r2", "m2", "kilo-model"),
				"\n",
			),
		),
		0644,
	)
	s.Restart()
	s.Tick()
	r, e := s.RESTClient(t).GetSessionDetailWithResponse(
		context.Background(),
		"echo",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.FatalNotNil(t, r.JSON200.Usage)
	unknown := map[string]bool{}

	for _, m := range *r.JSON200.Usage {
		unknown[m.Model] = m.UnknownModel
	}

	assert.Count(t, 2, unknown)
	assert.False(t, unknown["sonnet"])
	assert.True(t, unknown["kilo-model"])
}

func TestAnEmptyCostStillCarriesItsWindow(t *testing.T) {
	s := base.New(t)
	days := 3
	r, e := s.RESTClient(t).GetCostWithResponse(
		context.Background(),
		&client.GetCostParams{Days: &days},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 3, r.JSON200.Days)
}
