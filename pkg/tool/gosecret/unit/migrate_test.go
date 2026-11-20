package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/unit/secret_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestMigrateStoresManifestValues(t *testing.T) {
	s := secret_tester.New(t).Client()
	r, e := gosecret.Migrate(
		s,
		secret_tester.New(t).Manifest(),
		"alfa",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa/example-secret", r.Path)
	assert.Any(
		t,
		map[string]string{"ALFA": "one", "CHARLIE": "three"},
		s.Secrets["alfa/example-secret"],
	)
}

func TestMigrateTurnsManifestIntoExternalSecret(t *testing.T) {
	path := secret_tester.New(t).Manifest()
	_, e := gosecret.Migrate(
		secret_tester.New(t).Client(),
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

func TestMigrateTakesNamespaceFromDirectory(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "bravo")
	errors.PanicOnError(os.Mkdir(directory, 0o755))
	path := filepath.Join(directory, "example-secret.yaml")
	errors.PanicOnError(os.Rename(secret_tester.New(t).Manifest(), path))
	r, e := gosecret.Migrate(
		secret_tester.New(t).Client(),
		path,
		"",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo/example-secret", r.Path)
}

func TestMigrateRemovesDecodedFile(t *testing.T) {
	path := secret_tester.New(t).Manifest()
	decoded := gosecret.GetDecodedPath(path)
	system.SaveFile(decoded, "=== ALFA ===\none\n=== CHARLIE ===\nthree\n")
	r, e := gosecret.Migrate(
		secret_tester.New(t).Client(),
		path,
		"alfa",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.String(t, decoded, r.Removed)
	assert.False(t, system.FileExists(decoded))
}

func TestMigrateRefusesDecodedFileThatDiffers(t *testing.T) {
	path := secret_tester.New(t).Manifest()
	system.SaveFile(
		gosecret.GetDecodedPath(path),
		"=== ALFA ===\nedited\n=== CHARLIE ===\nthree\n",
	)
	s := secret_tester.New(t).Client()
	_, e := gosecret.Migrate(s, path, "alfa", constant.DefaultStore)
	assert.True(t, validation.Is(e))
	assert.Count(t, 0, s.Secrets["alfa/example-secret"])
}

func TestMigrateLeavesManifestWhenStoreDisagrees(t *testing.T) {
	path := secret_tester.New(t).Manifest()
	s := secret_tester.New(t).Client()
	s.Tamper = "ALFA"
	_, e := gosecret.Migrate(s, path, "alfa", constant.DefaultStore)
	assert.StringContains(t, "stored value differs", e.Error())
	assert.StringContains(
		t,
		"kind: Secret",
		system.ReadFile(filepath.Dir(path), filepath.Base(path)),
	)
}

func TestMigrateReportsKeysAlreadyInStore(t *testing.T) {
	s := secret_tester.New(t).Client()
	s.Secrets["alfa/example-secret"] = map[string]string{"DELTA": "four"}
	r, e := gosecret.Migrate(
		s,
		secret_tester.New(t).Manifest(),
		"alfa",
		constant.DefaultStore,
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"DELTA"}, r.Extra)
}

func TestExternalSecretCarriesNonOpaqueType(t *testing.T) {
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
