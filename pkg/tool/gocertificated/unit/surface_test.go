package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/publish_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecret"
	"net/url"
	"path/filepath"
	"testing"
)

func TestEveryToolIsRegistered(t *testing.T) {
	assert.Count(t, 12, base.New(t).Tools().ListTools())
}

func TestChainIsBuiltThroughTools(t *testing.T) {
	s := base.New(t)
	s.CreateRoot()
	assert.StringContains(t, "Example Issuing CA", s.CreateCluster())
	assert.StringContains(
		t,
		"Example",
		s.Tools().MustCallTool(constant.ListAuthorities, map[string]any{}),
	)
}

func TestSecondRootIsRefusedThroughTools(t *testing.T) {
	s := base.New(t)
	s.CreateRoot()
	assert.StringContains(t, "already live", s.CreateRootRefusal())
}

func TestIssuedCertificateReturnsItsKeyOnce(t *testing.T) {
	s := base.New(t)
	s.CreateRoot()
	s.CreateCluster()
	c := s.Tools()
	issued := c.MustCallTool(
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
		c.MustCallTool(constant.ListCertificates, map[string]any{}),
	)
}

func TestRootCertificateToolServesAnchor(t *testing.T) {
	s := base.New(t)
	s.CreateRoot()
	assert.StringContains(
		t,
		"BEGIN CERTIFICATE",
		s.Tools().MustCallTool(constant.RootCertificate, map[string]any{}),
	)
}

func TestPublishThroughToolsCommitsChain(t *testing.T) {
	s := base.New(t)
	s.CreateRoot()
	s.CreateCluster()
	c := s.Tools()
	c.MustCallTool(constant.Publish, map[string]any{})
	assert.Integer(t, 1, len(s.Forge.Commits()))
	assert.Integer(t, 6, len(s.Forge.Commits()[0].Actions))
	assert.StringContains(
		t,
		"Nothing to publish",
		c.MustCallTool(constant.Publish, map[string]any{}),
	)
}

func TestFirstPublishWritesBothRootFilesInOneCommit(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.Publish()
	commit := o.Forge.Commits()
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
	assert.String(t, "create", string(*o.Forge.Commits()[0].Actions[0].Action))
}

func TestExistingFileIsUpdatedNotCreated(t *testing.T) {
	o := publish_tester.New(t)
	o.Forge.SeedFile("certificate/root/certificate.pem", "stale")
	o.CreateRoot()
	o.Publish()
	assert.String(t, "update", string(*o.Forge.Commits()[0].Actions[0].Action))
}

func TestSecondPublishOnlyCarriesNewAuthority(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.Publish()
	o.CreateCluster()
	o.Publish()
	commit := o.Forge.Commits()
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
	assert.Integer(t, 1, len(o.Forge.Commits()))
}

func TestLeafCertificatesAreNeverPublished(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.IssueLeaf()
	o.Publish()
	commit := o.Forge.Commits()
	assert.Integer(t, 1, len(commit))
	assert.Integer(t, 6, len(commit[0].Actions))
}

func TestOnlyNamedAuthorityBecomesSecret(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.Publish()
	assert.Strings(
		t,
		[]string{
			"certificate/root/certificate.pem",
			"certificate/root/key.pem",
		},
		publish_tester.Paths(o.Forge.Commits()[0].Actions),
	)
}

func TestClusterAuthorityDeliversSecretPair(t *testing.T) {
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
		publish_tester.Paths(o.Forge.Commits()[0].Actions),
	)
}

func TestDeliveredSecretIsTlsManifest(t *testing.T) {
	o := publish_tester.New(t)
	o.CreateRoot()
	o.CreateCluster()
	o.Publish()
	manifest := publish_tester.Content(
		o.Forge.Commits()[0].Actions,
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
	action := o.Forge.Commits()[0].Actions
	directory := t.TempDir()
	manifest := filepath.Join(directory, "authority-secret.yaml")
	system_tester.New(t).Write(
		directory,
		"authority-secret.yaml",
		publish_tester.Content(action, constant.FixtureSecretPath),
	)
	system_tester.New(t).Write(
		directory,
		gosecret.GetDecodedPath("authority-secret.yaml"),
		publish_tester.Content(
			action,
			gosecret.GetDecodedPath(constant.FixtureSecretPath),
		),
	)
	result, e := gosecret.EncodeSecret(manifest)
	assert.Nil(t, e)
	assert.True(t, result.InSync)
}

func TestEmptyDashboardInvitesRoot(t *testing.T) {
	assert.StringContains(
		t,
		"Create the root to begin the chain",
		base.New(t).Page(constant.DashboardPath),
	)
}

func TestEveryPageRenders(t *testing.T) {
	s := base.New(t)
	s.Page(constant.DashboardPath)
	s.Page(constant.AuthoritiesPath)
	s.Page(constant.CertificatesPath)
	s.Page(constant.CreateAuthorityPath)
	s.Page(constant.IssueCertificatePath)
}

func TestRootIsCreatedThroughForm(t *testing.T) {
	s := base.New(t)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	assert.StringContains(
		t,
		"Example Root CA",
		s.Page(constant.AuthoritiesPath),
	)
}

func TestFormRejectionKeepsValues(t *testing.T) {
	s := base.New(t)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	page := s.Submit(constant.CreateAuthorityPath, base.RootValues())
	assert.StringContains(t, "already live", page)
	assert.StringContains(t, "Example Root CA", page)
}

func TestIssuedKeyIsShownOnceThenNeverAgain(t *testing.T) {
	s := base.New(t)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	s.Submit(constant.CreateAuthorityPath, base.ClusterValues())
	assert.StringContains(
		t,
		"PRIVATE KEY",
		s.Submit(constant.IssueCertificatePath, base.LeafValues()),
	)
	assert.StringNotContains(
		t,
		"PRIVATE KEY",
		s.Page(constant.CertificatesPath),
	)
}

func TestPublishButtonCommitsChain(t *testing.T) {
	s := base.New(t)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	assert.StringContains(
		t,
		"certificate/root/certificate.pem",
		s.Page(constant.DashboardPath),
	)
	s.Submit(constant.PublishPath, url.Values{})
	assert.Integer(t, 1, len(s.Forge.Commits()))
	assert.StringContains(
		t,
		"Everything is published",
		s.Page(constant.DashboardPath),
	)
}

func TestRootPathServesAnchorAsText(t *testing.T) {
	s := base.New(t)
	s.Submit(constant.CreateAuthorityPath, base.RootValues())
	assert.StringContains(t, "BEGIN CERTIFICATE", s.Page(constant.RootPath))
}
