package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/authority_tester"
	"testing"
)

func TestRootIsSelfSignedAuthority(t *testing.T) {
	c := authority_tester.NewRoot().Material().Certificate
	assert.True(t, c.IsCA)
	assert.String(t, c.Subject.String(), c.Issuer.String())
	assert.String(t, "Example Root CA", c.Subject.CommonName)
}

func TestIntermediateInheritsIssuerOrganization(t *testing.T) {
	c := authority_tester.NewCluster(authority_tester.NewRoot()).Material().Certificate
	assert.String(t, "Example", c.Subject.Organization[0])
	assert.String(t, "XX", c.Subject.Country[0])
}

func TestIntermediateForbidsFurtherAuthority(t *testing.T) {
	c := authority_tester.NewCluster(authority_tester.NewRoot()).Material().Certificate
	assert.True(t, c.IsCA)
	assert.True(t, c.MaxPathLenZero)
	assert.Integer(t, 0, c.MaxPathLen)
}

func TestPermittedDomainVerifies(t *testing.T) {
	root := authority_tester.NewRoot()
	cluster := authority_tester.NewCluster(root)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, authority_tester.Verify(root, cluster, leaf.Certificate))
}

func TestPermittedLocalDomainVerifies(t *testing.T) {
	root := authority_tester.NewRoot()
	cluster := authority_tester.NewCluster(root)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureLocalHost},
	)
	assert.Nil(t, authority_tester.Verify(root, cluster, leaf.Certificate))
}

func TestForeignDomainFailsVerification(t *testing.T) {
	root := authority_tester.NewRoot()
	cluster := authority_tester.NewCluster(root)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureImpostor,
		[]string{constant.FixtureForeignDomain},
	)
	assert.Error(t, authority_tester.Verify(root, cluster, leaf.Certificate))
}

func TestForeignAddressFailsVerification(t *testing.T) {
	root := authority_tester.NewRoot()
	cluster := authority_tester.NewCluster(root)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureImpostor,
		[]string{constant.FixtureForeignHost},
	)
	assert.Error(t, authority_tester.Verify(root, cluster, leaf.Certificate))
}

func TestPermittedAddressVerifies(t *testing.T) {
	root := authority_tester.NewRoot()
	cluster := authority_tester.NewCluster(root)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixturePermittedHost},
	)
	assert.Nil(t, authority_tester.Verify(root, cluster, leaf.Certificate))
}
