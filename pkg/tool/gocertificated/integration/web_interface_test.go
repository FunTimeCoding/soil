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
	defer o.Server.Close()
	assert.StringContains(
		t,
		"Create the root to begin the chain",
		o.Page(t, constant.DashboardPath),
	)
}

func TestEveryPageRenders(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Server.Close()
	o.Page(t, constant.DashboardPath)
	o.Page(t, constant.AuthoritiesPath)
	o.Page(t, constant.CertificatesPath)
	o.Page(t, constant.CreateAuthorityPath)
	o.Page(t, constant.IssueCertificatePath)
}

func TestRootIsCreatedThroughTheForm(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Server.Close()
	o.Submit(t, constant.CreateAuthorityPath, rootValues())
	assert.StringContains(
		t,
		"Example Root CA",
		o.Page(t, constant.AuthoritiesPath),
	)
}

func TestFormRejectionKeepsTheValues(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Server.Close()
	o.Submit(t, constant.CreateAuthorityPath, rootValues())
	page := o.Submit(t, constant.CreateAuthorityPath, rootValues())
	assert.StringContains(t, "already live", page)
	assert.StringContains(t, "Example Root CA", page)
}

func TestIssuedKeyIsShownOnceThenNeverAgain(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Server.Close()
	o.Submit(t, constant.CreateAuthorityPath, rootValues())
	o.Submit(t, constant.CreateAuthorityPath, clusterValues())
	assert.StringContains(
		t,
		"PRIVATE KEY",
		o.Submit(t, constant.IssueCertificatePath, leafValues()),
	)
	assert.StringNotContains(
		t,
		"PRIVATE KEY",
		o.Page(t, constant.CertificatesPath),
	)
}

func TestPublishButtonCommitsTheChain(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Server.Close()
	o.Submit(t, constant.CreateAuthorityPath, rootValues())
	assert.StringContains(
		t,
		"certificate/root/certificate.pem",
		o.Page(t, constant.DashboardPath),
	)
	o.Submit(t, constant.PublishPath, url.Values{})
	assert.Integer(t, 1, len(o.Server.Forge.Commits()))
	assert.StringContains(
		t,
		"Everything is published",
		o.Page(t, constant.DashboardPath),
	)
}

func TestRootPathServesTheAnchorAsText(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Server.Close()
	o.Submit(t, constant.CreateAuthorityPath, rootValues())
	assert.StringContains(t, "BEGIN CERTIFICATE", o.Page(t, constant.RootPath))
}

func rootValues() url.Values {
	return url.Values{
		"name":                         {constant.RootAuthority},
		constant.KindParameter:         {string(constant.KindRoot)},
		constant.CommonNameParameter:   {constant.FixtureRootCommonName},
		constant.CountryParameter:      {constant.FixtureCountry},
		constant.ProvinceParameter:     {constant.FixtureProvince},
		constant.OrganizationParameter: {constant.FixtureOrganization},
	}
}

func clusterValues() url.Values {
	return url.Values{
		"name":                       {constant.FixtureClusterAuthority},
		constant.KindParameter:       {string(constant.KindIntermediate)},
		constant.CommonNameParameter: {constant.FixtureIssuingCommonName},
		constant.DomainParameter:     {constant.FixtureDomain},
		constant.AddressParameter:    {constant.FixtureAddress},
	}
}

func leafValues() url.Values {
	return url.Values{
		constant.AuthorityParameter:  {constant.FixtureClusterAuthority},
		constant.KindParameter:       {string(constant.KindServer)},
		constant.CommonNameParameter: {constant.FixtureCommonName},
		constant.HostParameter:       {constant.FixtureHost},
	}
}
