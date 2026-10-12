package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/unit/store_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMigrateStoresTheManifestValues(t *testing.T) {
	s := store_tester.New()
	r, e := gosecret.Migrate(s, writeSecret(t), "alfa", constant.DefaultStore)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa/example-secret", r.Path)
	assert.Any(
		t,
		map[string]string{"ALFA": "one", "CHARLIE": "three"},
		s.Stored("alfa/example-secret"),
	)
}

func TestMigrateTurnsTheManifestIntoAnExternalSecret(t *testing.T) {
	path := writeSecret(t)
	_, e := gosecret.Migrate(
		store_tester.New(),
		path,
		"alfa",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`---
apiVersion: external-secrets.io/v1
kind: ExternalSecret
metadata: {name: example-secret}
spec:
  refreshInterval: 1h
  secretStoreRef: {kind: ClusterSecretStore, name: openbao}
  target: {name: example-secret, creationPolicy: Owner}
  dataFrom:
  - extract: {key: alfa/example-secret}
`,
		system.ReadFile(filepath.Dir(path), filepath.Base(path)),
	)
}

func TestMigrateTakesTheNamespaceFromTheDirectory(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "bravo")
	errors.PanicOnError(os.Mkdir(directory, 0o755))
	path := filepath.Join(directory, "example-secret.yaml")
	errors.PanicOnError(os.Rename(writeSecret(t), path))
	r, e := gosecret.Migrate(
		store_tester.New(),
		path,
		"",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo/example-secret", r.Path)
}

func TestMigrateRemovesTheDecodedFile(t *testing.T) {
	path := writeSecret(t)
	decoded := gosecret.GetDecodedPath(path)
	system.SaveFile(decoded, "=== ALFA ===\none\n=== CHARLIE ===\nthree\n")
	r, e := gosecret.Migrate(
		store_tester.New(),
		path,
		"alfa",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.String(t, decoded, r.Removed)
	assert.False(t, system.FileExists(decoded))
}

func TestMigrateRefusesADecodedFileThatDiffers(t *testing.T) {
	path := writeSecret(t)
	system.SaveFile(
		gosecret.GetDecodedPath(path),
		"=== ALFA ===\nedited\n=== CHARLIE ===\nthree\n",
	)
	s := store_tester.New()
	_, e := gosecret.Migrate(s, path, "alfa", constant.DefaultStore)
	assert.True(t, validation.Is(e))
	assert.Count(t, 0, s.Stored("alfa/example-secret"))
}

func TestMigrateLeavesTheManifestWhenTheStoreDisagrees(t *testing.T) {
	path := writeSecret(t)
	_, e := gosecret.Migrate(
		store_tester.New().WithTamper("ALFA"),
		path,
		"alfa",
		constant.DefaultStore,
	)
	assert.StringContains(t, "stored value differs", e.Error())
	assert.StringContains(
		t,
		"kind: Secret",
		system.ReadFile(filepath.Dir(path), filepath.Base(path)),
	)
}

func TestMigrateReportsKeysAlreadyInTheStore(t *testing.T) {
	r, e := gosecret.Migrate(
		store_tester.New().WithExisting(map[string]string{"DELTA": "four"}),
		writeSecret(t),
		"alfa",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"DELTA"}, r.Extra)
}

func TestExternalSecretCarriesANonOpaqueType(t *testing.T) {
	assert.StringContains(
		t,
		"creationPolicy: Owner, template: {type: kubernetes.io/tls}}",
		gosecret.ExternalSecret(
			"alfa-fsec",
			"alfa/alfa-fsec",
			constant.DefaultStore,
			"kubernetes.io/tls",
		),
	)
}
