package unit

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/key"
	"testing"
)

func TestKeyEncodeJoinsEveryFieldOfEveryRow(t *testing.T) {
	assert.String(
		t,
		"alfa|YWxmYQ==|hmac-sha256|bravo|YnJhdm8=|hmac-sha512",
		key.Encode(
			settings(
				`{"tsigKeys":[
					{"keyName":"alfa","sharedSecret":"YWxmYQ==","algorithmName":"hmac-sha256"},
					{"keyName":"bravo","sharedSecret":"YnJhdm8=","algorithmName":"hmac-sha512"}
				]}`,
			).Keys,
		),
	)
}

func TestKeyEncodeOfNothingClearsTheList(t *testing.T) {
	assert.String(t, "false", key.Encode(nil))
}

func TestKeyGenerateMakesA32ByteSecret(t *testing.T) {
	k := key.Generate("alfa", constant.DefaultAlgorithm)
	b, e := base64.StdEncoding.DecodeString(k.Secret)
	assert.FatalOnError(t, e)
	assert.Count(t, 32, b)
	assert.String(t, "alfa", k.Name)
	assert.String(t, "hmac-sha256", k.Algorithm)
}

func TestKeyGenerateNeverRepeatsASecret(t *testing.T) {
	assert.True(
		t,
		key.Generate("alfa", constant.DefaultAlgorithm).Secret != key.Generate(
			"alfa",
			constant.DefaultAlgorithm,
		).Secret,
	)
}

func TestKeyFindByName(t *testing.T) {
	v := settings(
		`{"tsigKeys":[{"keyName":"alfa","sharedSecret":"YWxmYQ==","algorithmName":"hmac-sha256"}]}`,
	).Keys
	assert.String(t, "YWxmYQ==", key.Find(v, "alfa").Secret)
	assert.True(t, key.Find(v, "bravo") == nil)
}
