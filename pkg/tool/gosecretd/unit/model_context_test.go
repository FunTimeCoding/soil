package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/model_context_tester"
	"testing"
)

func TestModelContextOffersFiveTools(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.Count(t, 5, o.Client.ListTools())
}

func TestStatusReportsTheSeal(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Server.MockClient.Seal()
	assert.StringContains(
		t,
		`"sealed": true`,
		o.Client.MustCallTool(constant.Status, map[string]any{}),
	)
}

func TestShowSecretMasksValues(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Server.MockClient.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value"},
	)
	result := o.Client.MustCallTool(
		constant.ShowSecret,
		map[string]any{"path": "alfa/bravo"},
	)
	assert.StringContains(t, "CHARLIE", result)
	assert.False(t, contains(result, "delta-value"))
}

func TestRevealSecretReturnsOneValue(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Server.MockClient.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value", "ECHO": "foxtrot-value"},
	)
	result := o.Client.MustCallTool(
		constant.RevealSecret,
		map[string]any{"path": "alfa/bravo", "key": "CHARLIE"},
	)
	assert.StringContains(t, "delta-value", result)
	assert.False(t, contains(result, "foxtrot-value"))
}

func TestRevealOfAMissingKeyFailsWithoutACapture(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Server.MockClient.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value"},
	)
	assert.StringContains(
		t,
		"secret key not found: GOLF",
		o.Client.MustCallToolError(
			constant.RevealSecret,
			map[string]any{"path": "alfa/bravo", "key": "GOLF"},
		),
	)
	assert.Count(t, 0, o.Server.Reporter.Events())
}

func TestPutSecretMergesKeys(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Server.MockClient.AddSecret(
		"alfa/bravo",
		map[string]string{"CHARLIE": "delta-value"},
	)
	o.Client.MustCallTool(
		constant.PutSecret,
		map[string]any{
			"path":   "alfa/bravo",
			"values": map[string]any{"ECHO": "foxtrot-value"},
		},
	)
	assert.Any(
		t,
		map[string]string{"CHARLIE": "delta-value", "ECHO": "foxtrot-value"},
		o.Server.MockClient.Stored("alfa/bravo"),
	)
}

func TestPathClimbingOutIsRefused(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.StringContains(
		t,
		"path climbs out of the store",
		o.Client.MustCallToolError(
			constant.ShowSecret,
			map[string]any{"path": "alfa/../bravo"},
		),
	)
}

func TestListSecretsShowsFolders(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Server.MockClient.AddSecret("alfa/bravo", map[string]string{"A": "1"})
	o.Server.MockClient.AddSecret(
		"alfa/charlie/delta",
		map[string]string{"B": "2"},
	)
	result := o.Client.MustCallTool(
		constant.ListSecrets,
		map[string]any{"path": "alfa"},
	)
	assert.StringContains(t, `"bravo"`, result)
	assert.StringContains(t, `"charlie/"`, result)
}
