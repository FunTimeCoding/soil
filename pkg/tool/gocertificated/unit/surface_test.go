package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/model_context_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/publish_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/web_interface_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"net/url"
	"path/filepath"
	"testing"
)

func TestEveryToolIsRegistered(t *testing.T) {
	o := model_context_tester.New(t)
	assert.Count(t, 12, o.Client.ListTools())
}

func TestChainIsBuiltThroughTools(t *testing.T) {
	o := model_context_tester.New(t)
	o.CreateRoot()
	assert.StringContains(t, "Example Issuing CA", o.CreateCluster())
	assert.StringContains(
		t,
		"Example",
		o.Client.MustCallTool(constant.ListAuthorities, map[string]any{}),
	)
}

func TestSecondRootIsRefusedThroughTools(t *testing.T) {
	o := model_context_tester.New(t)
	o.CreateRoot()
	assert.StringContains(
		t,
		"An authority of that name is already live",
		o.CreateRootRefusal(),
	)
}

func TestIssuedCertificateReturnsItsKeyOnce(t *testing.T) {
	o := model_context_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	issued := o.Client.MustCallTool(
		constant.IssueCertificate,
		map[string]any{
			constant.AuthorityParameter:  constant.FixtureClusterAuthority,
			constant.KindParameter:       string(constant.KindServer),
			constant.CommonNameParameter: constant.FixtureCommonName,
			constant.HostParameter:       []any{constant.FixtureHost},
		},
	)
	assert.StringContains(t, "PRIVATE KEY", issued)
	assert.StringNotContains(
		t,
		"PRIVATE KEY",
		o.Client.MustCallTool(constant.ListCertificates, map[string]any{}),
	)
}

func TestRootCertificateToolServesTheAnchor(t *testing.T) {
	o := model_context_tester.New(t)
	o.CreateRoot()
	assert.StringContains(
		t,
		"BEGIN CERTIFICATE",
		o.Client.MustCallTool(constant.RootCertificate, map[string]any{}),
	)
}

func TestPublishThroughToolsCommitsTheChain(t *testing.T) {
	o := model_context_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.Client.MustCallTool(constant.Publish, map[string]any{})
	assert.Integer(t, 1, len(o.Server.Forge.Commits()))
	assert.Integer(t, 6, len(o.Server.Forge.Commits()[0].Actions))
	assert.StringContains(
		t,
		"Nothing to publish",
		o.Client.MustCallTool(constant.Publish, map[string]any{}),
	)
}

func TestFirstPublishWritesBothRootFilesInOneCommit(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.Publish()
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
	o.CreateRoot()
	o.Publish()
	assert.String(
		t,
		"create",
		string(*o.Server.Forge.Commits()[0].Actions[0].Action),
	)
}

func TestExistingFileIsUpdatedNotCreated(t *testing.T) {
	o := publish_tester.New(t)
	o.Server.Forge.SeedFile("certificate/root/certificate.pem", "stale")
	o.CreateRoot()
	o.Publish()
	assert.String(
		t,
		"update",
		string(*o.Server.Forge.Commits()[0].Actions[0].Action),
	)
}

func TestSecondPublishOnlyCarriesTheNewAuthority(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.Publish()
	o.CreateCluster()
	o.Publish()
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
	o.CreateRoot()
	o.Publish()
	o.Publish()
	assert.Integer(t, 1, len(o.Server.Forge.Commits()))
}

func TestLeafCertificatesAreNeverPublished(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.IssueLeaf()
	o.Publish()
	commit := o.Server.Forge.Commits()
	assert.Integer(t, 1, len(commit))
	assert.Integer(t, 6, len(commit[0].Actions))
}

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

func TestEmptyDashboardInvitesTheRoot(t *testing.T) {
	o := web_interface_tester.New(t)
	assert.StringContains(
		t,
		"Create the root to begin the chain",
		o.Page(constant.DashboardPath),
	)
}

func TestEveryPageRenders(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Page(constant.DashboardPath)
	o.Page(constant.AuthoritiesPath)
	o.Page(constant.CertificatesPath)
	o.Page(constant.CreateAuthorityPath)
	o.Page(constant.IssueCertificatePath)
}

func TestRootIsCreatedThroughTheForm(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Submit(constant.CreateAuthorityPath, rootValues())
	assert.StringContains(
		t,
		"Example Root CA",
		o.Page(constant.AuthoritiesPath),
	)
}

func TestFormRejectionKeepsTheValues(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Submit(constant.CreateAuthorityPath, rootValues())
	page := o.Submit(constant.CreateAuthorityPath, rootValues())
	assert.StringContains(t, "already live", page)
	assert.StringContains(t, "Example Root CA", page)
}

func TestIssuedKeyIsShownOnceThenNeverAgain(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Submit(constant.CreateAuthorityPath, rootValues())
	o.Submit(constant.CreateAuthorityPath, clusterValues())
	assert.StringContains(
		t,
		"PRIVATE KEY",
		o.Submit(constant.IssueCertificatePath, leafValues()),
	)
	assert.StringNotContains(
		t,
		"PRIVATE KEY",
		o.Page(constant.CertificatesPath),
	)
}

func TestPublishButtonCommitsTheChain(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Submit(constant.CreateAuthorityPath, rootValues())
	assert.StringContains(
		t,
		"certificate/root/certificate.pem",
		o.Page(constant.DashboardPath),
	)
	o.Submit(constant.PublishPath, url.Values{})
	assert.Integer(t, 1, len(o.Server.Forge.Commits()))
	assert.StringContains(
		t,
		"Everything is published",
		o.Page(constant.DashboardPath),
	)
}

func TestRootPathServesTheAnchorAsText(t *testing.T) {
	o := web_interface_tester.New(t)
	o.Submit(constant.CreateAuthorityPath, rootValues())
	assert.StringContains(t, "BEGIN CERTIFICATE", o.Page(constant.RootPath))
}
