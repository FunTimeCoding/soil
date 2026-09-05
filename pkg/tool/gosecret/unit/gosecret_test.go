package unit

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeSecret(t *testing.T) string {
	path := filepath.Join(t.TempDir(), "example-secret.yaml")
	errors.PanicOnError(os.WriteFile(path, []byte(constant.TestManifest), 0644))

	return path
}

func TestEncodeSecret(t *testing.T) {
	path := writeSecret(t)
	decodedPath := gosecret.GetDecodedPath(path)
	errors.PanicOnError(
		os.WriteFile(
			decodedPath,
			[]byte("=== ALPHA ===\ntwo\n=== BETA ===\nfirst\nsecond\n"),
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
	alpha, e := base64.StdEncoding.DecodeString(m.Payload["ALPHA"])
	errors.PanicOnError(e)
	assert.String(t, "two", string(alpha))
	beta, e := base64.StdEncoding.DecodeString(m.Payload["BETA"])
	errors.PanicOnError(e)
	assert.String(t, "first\nsecond\n", string(beta))
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
		map[string]string{"ALPHA": "one", "GAMMA": "three"},
	)
	errors.PanicOnError(e)
	assert.True(t, inSync)
}
