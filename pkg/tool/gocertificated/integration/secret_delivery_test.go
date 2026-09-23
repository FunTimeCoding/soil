package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration/publish_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"path/filepath"
	"testing"
)

func TestOnlyTheNamedAuthorityBecomesASecret(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.Publish()
	assert.Strings(
		t,
		[]string{
			"certificate/root/certificate.pem",
			"certificate/root/key.pem",
		},
		publish_tester.Paths(o.Server.Forge.Commits()[0].Actions),
	)
}

func TestClusterAuthorityDeliversTheSecretPair(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.Publish()
	assert.Strings(
		t,
		[]string{
			"certificate/root/certificate.pem",
			"certificate/root/key.pem",
			"certificate/cluster/certificate.pem",
			"certificate/cluster/key.pem",
			"manifest/authority-secret.yaml",
			"manifest/authority-secret.decoded.txt",
		},
		publish_tester.Paths(o.Server.Forge.Commits()[0].Actions),
	)
}

func TestDeliveredSecretIsATlsManifest(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.Publish()
	manifest := publish_tester.Content(
		o.Server.Forge.Commits()[0].Actions,
		constant.FixtureSecretPath,
	)
	assert.StringContains(t, "kubernetes.io/tls", manifest)
	assert.StringContains(t, "name: authority-secret", manifest)
	assert.StringContains(t, "tls.crt:", manifest)
	assert.StringContains(t, "tls.key:", manifest)
}

func TestDeliveredPairSatisfiesGosecret(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.Publish()
	action := o.Server.Forge.Commits()[0].Actions
	directory := t.TempDir()
	manifest := filepath.Join(directory, "authority-secret.yaml")
	write(
		t,
		manifest,
		publish_tester.Content(action, constant.FixtureSecretPath),
	)
	write(
		t,
		gosecret.GetDecodedPath(manifest),
		publish_tester.Content(
			action,
			gosecret.GetDecodedPath(constant.FixtureSecretPath),
		),
	)
	result, e := gosecret.EncodeSecret(manifest)
	assert.Nil(t, e)
	assert.True(t, result.InSync)
}
