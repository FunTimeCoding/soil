package unit

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"go.yaml.in/yaml/v3"
	"os"
	"strings"
	"testing"
)

func TestEncodeSecret(t *testing.T) {
	path := writeSecret(t)
	decodedPath := gosecret.GetDecodedPath(path)
	errors.PanicOnError(
		os.WriteFile(
			decodedPath,
			[]byte("=== ALFA ===\ntwo\n=== BRAVO ===\nfirst\nsecond\n"),
			0600,
		),
	)
	s, e := gosecret.EncodeSecret(path)
	errors.PanicOnError(e)
	assert.True(t, !s.InSync)
	b, e := os.ReadFile(path)
	errors.PanicOnError(e)
	content := string(b)
	assert.True(t, strings.Contains(content, "# noinspection"))
	assert.True(t, strings.HasPrefix(content, "---\napiVersion: v1"))
	var m gosecret.SecretManifest
	errors.PanicOnError(yaml.Unmarshal(b, &m))
	assert.Integer(t, 2, len(m.Payload))
	alfa, e := base64.StdEncoding.DecodeString(m.Payload["ALFA"])
	errors.PanicOnError(e)
	assert.String(t, "two", string(alfa))
	bravo, e := base64.StdEncoding.DecodeString(m.Payload["BRAVO"])
	errors.PanicOnError(e)
	assert.String(t, "first\nsecond\n", string(bravo))
	s, e = gosecret.EncodeSecret(path)
	errors.PanicOnError(e)
	assert.True(t, s.InSync)
}

func TestEncodeSecretWithoutDecoded(t *testing.T) {
	s, e := gosecret.EncodeSecret(writeSecret(t))
	errors.PanicOnError(e)
	assert.True(t, s == nil)
}

func TestCheckSyncAfterDecode(t *testing.T) {
	path := writeSecret(t)
	s, e := gosecret.ProcessSecret(path, false)
	errors.PanicOnError(e)
	assert.True(t, s.InSync)
	inSync, e := gosecret.CheckSync(
		s.DecodedPath,
		map[string]string{"ALFA": "one", "CHARLIE": "three"},
	)
	errors.PanicOnError(e)
	assert.True(t, inSync)
}
