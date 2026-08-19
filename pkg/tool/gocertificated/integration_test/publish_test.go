package integration_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration_test/publish_tester"
	"testing"
)

func TestFirstPublishWritesBothRootFilesInOneCommit(t *testing.T) {
	o := publish_tester.New(t)
	defer o.Server.Close()
	o.CreateRoot(t)
	o.Publish(t)
	commit := o.Server.Forge.Commits()
	assert.Integer(t, 1, len(commit))
	assert.Strings(
		t,
		[]string{
			"certificate/root/certificate.pem",
			"certificate/root/key.pem",
		},
		publish_tester.Paths(commit[0].Actions),
	)
}

func TestMissingFileIsCreatedNotUpdated(t *testing.T) {
	o := publish_tester.New(t)
	defer o.Server.Close()
	o.CreateRoot(t)
	o.Publish(t)
	assert.String(
		t,
		"create",
		string(*o.Server.Forge.Commits()[0].Actions[0].Action),
	)
}

func TestExistingFileIsUpdatedNotCreated(t *testing.T) {
	o := publish_tester.New(t)
	defer o.Server.Close()
	o.Server.Forge.SeedFile("certificate/root/certificate.pem", "stale")
	o.CreateRoot(t)
	o.Publish(t)
	assert.String(
		t,
		"update",
		string(*o.Server.Forge.Commits()[0].Actions[0].Action),
	)
}

func TestSecondPublishOnlyCarriesTheNewAuthority(t *testing.T) {
	o := publish_tester.New(t)
	defer o.Server.Close()
	o.CreateRoot(t)
	o.Publish(t)
	o.CreateCluster(t)
	o.Publish(t)
	commit := o.Server.Forge.Commits()
	assert.Integer(t, 2, len(commit))
	assert.Strings(
		t,
		[]string{
			"certificate/cluster/certificate.pem",
			"certificate/cluster/key.pem",
			"manifest/authority-secret.yaml",
			"manifest/authority-secret.decoded.txt",
		},
		publish_tester.Paths(commit[1].Actions),
	)
}

func TestPublishWithNothingPendingWritesNoCommit(t *testing.T) {
	o := publish_tester.New(t)
	defer o.Server.Close()
	o.CreateRoot(t)
	o.Publish(t)
	o.Publish(t)
	assert.Integer(t, 1, len(o.Server.Forge.Commits()))
}

func TestLeafCertificatesAreNeverPublished(t *testing.T) {
	o := publish_tester.New(t)
	defer o.Server.Close()
	o.CreateRoot(t)
	o.CreateCluster(t)
	o.IssueLeaf(t)
	o.Publish(t)
	commit := o.Server.Forge.Commits()
	assert.Integer(t, 1, len(commit))
	assert.Integer(t, 6, len(commit[0].Actions))
}
