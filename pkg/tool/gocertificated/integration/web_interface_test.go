package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration/web_interface_tester"
	"net/url"
	"testing"
)

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
