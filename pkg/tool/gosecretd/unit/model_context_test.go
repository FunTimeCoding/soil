package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"testing"
)

func TestModelContextOffersFiveTools(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.Count(t, 5, c.ListTools())
}

func TestStatusReportsSeal(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Source.Seal()
	assert.StringContains(
		t,
		`"sealed": true`,
		c.MustCallTool(constant.Status, map[string]any{}),
	)
}

func TestShowSecretMasksValues(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Source.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value"},
	)
	result := c.MustCallTool(
		constant.ShowSecret,
		map[string]any{"path": "alfa/bravo"},
	)
	assert.StringContains(t, "CHARLIE", result)
	assert.StringNotContains(t, "delta-value", result)
}

func TestRevealSecretReturnsOneValue(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Source.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value", "ECHO": "foxtrot-value"},
	)
	result := c.MustCallTool(
		constant.RevealSecret,
		map[string]any{"path": "alfa/bravo", "key": "CHARLIE"},
	)
	assert.StringContains(t, "delta-value", result)
	assert.StringNotContains(t, "foxtrot-value", result)
}

func TestRevealOfMissingKeyFailsWithoutCapture(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Source.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value"},
	)
	assert.StringContains(
		t,
		"secret key not found: GOLF",
		c.MustCallToolError(
			constant.RevealSecret,
			map[string]any{"path": "alfa/bravo", "key": "GOLF"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestPutSecretMergesKeys(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Source.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value"},
	)
	c.MustCallTool(
		constant.PutSecret,
		map[string]any{
			"path":   "alfa/bravo",
			"values": map[string]any{"ECHO": "foxtrot-value"},
		},
	)
	assert.Any(
		t,
		map[string]string{"CHARLIE": "delta-value", "ECHO": "foxtrot-value"},
		s.Source.Stored("alfa/bravo"),
	)
}

func TestPathClimbingOutIsRefused(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.StringContains(
		t,
		"path climbs out of the store",
		c.MustCallToolError(
			constant.ShowSecret,
			map[string]any{"path": "alfa/../bravo"},
		),
	)
}

func TestListSecretsShowsFolders(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	s.Source.AddSecret("alfa/bravo", map[string]string{"A": "1"})
	s.Source.AddSecret("alfa/charlie/delta", map[string]string{"B": "2"})
	result := c.MustCallTool(
		constant.ListSecrets,
		map[string]any{"path": "alfa"},
	)
	assert.StringContains(t, `"bravo"`, result)
	assert.StringContains(t, `"charlie/"`, result)
}
