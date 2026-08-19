package unit_test

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/name_constraint"
	"net"
	"testing"
)

func newRootName() *distinguished_name.Name {
	n := distinguished_name.New()
	n.Country = constant.FixtureCountry
	n.Province = constant.FixtureProvince
	n.Organization = constant.FixtureOrganization
	n.CommonName = constant.FixtureRootCommonName

	return n
}

func newRoot() *authority.Authority {
	return authority.NewRoot(newRootName(), constant.RootValidityYear)
}

func newClusterConstraint() *name_constraint.Constraint {
	_, permitted, e := net.ParseCIDR(constant.FixtureAddress)

	if e != nil {
		panic(e)
	}

	c := name_constraint.New()
	c.PermittedDomain = []string{
		constant.FixtureDomain,
		constant.FixtureInternalDomain,
		constant.FixtureLocalDomain,
	}
	c.PermittedAddress = []*net.IPNet{permitted}

	return c
}

func newCluster(root *authority.Authority) *authority.Authority {
	n := distinguished_name.New()
	n.CommonName = constant.FixtureIssuingCommonName
	r := issue_request.New()
	r.Kind = constant.KindIntermediate
	r.Name = n
	r.Constraint = newClusterConstraint()
	r.ValidYear = constant.IntermediateValidityYear

	return authority.New(root.Issue(r))
}

func newLeaf(
	cluster *authority.Authority,
	common string,
	host []string,
) *material.Material {
	n := distinguished_name.New()
	n.CommonName = common
	r := issue_request.New()
	r.Kind = constant.KindServer
	r.Name = n
	r.Host = host
	r.ValidDay = constant.LeafValidityDay

	return cluster.Issue(r)
}

func verify(
	root *authority.Authority,
	cluster *authority.Authority,
	leaf *x509.Certificate,
) error {
	roots := x509.NewCertPool()
	roots.AddCert(root.Material().Certificate)
	middle := x509.NewCertPool()
	middle.AddCert(cluster.Material().Certificate)
	o := x509.VerifyOptions{}
	o.Roots = roots
	o.Intermediates = middle
	o.KeyUsages = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	_, e := leaf.Verify(o)

	return e
}

func TestRootIsSelfSignedAuthority(t *testing.T) {
	c := newRoot().Material().Certificate
	assert.True(t, c.IsCA)
	assert.String(t, c.Subject.String(), c.Issuer.String())
	assert.String(t, "Example Root CA", c.Subject.CommonName)
}

func TestIntermediateInheritsIssuerOrganization(t *testing.T) {
	c := newCluster(newRoot()).Material().Certificate
	assert.String(t, "Example", c.Subject.Organization[0])
	assert.String(t, "XX", c.Subject.Country[0])
}

func TestIntermediateForbidsFurtherAuthority(t *testing.T) {
	c := newCluster(newRoot()).Material().Certificate
	assert.True(t, c.IsCA)
	assert.True(t, c.MaxPathLenZero)
	assert.Integer(t, 0, c.MaxPathLen)
}

func TestPermittedDomainVerifies(t *testing.T) {
	root := newRoot()
	cluster := newCluster(root)
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, verify(root, cluster, leaf.Certificate))
}

func TestPermittedLocalDomainVerifies(t *testing.T) {
	root := newRoot()
	cluster := newCluster(root)
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureLocalHost},
	)
	assert.Nil(t, verify(root, cluster, leaf.Certificate))
}

func TestForeignDomainFailsVerification(t *testing.T) {
	root := newRoot()
	cluster := newCluster(root)
	leaf := newLeaf(
		cluster,
		constant.FixtureImpostor,
		[]string{constant.FixtureForeignDomain},
	)
	assert.Error(t, verify(root, cluster, leaf.Certificate))
}

func TestForeignAddressFailsVerification(t *testing.T) {
	root := newRoot()
	cluster := newCluster(root)
	leaf := newLeaf(
		cluster,
		constant.FixtureImpostor,
		[]string{constant.FixtureForeignHost},
	)
	assert.Error(t, verify(root, cluster, leaf.Certificate))
}

func TestPermittedAddressVerifies(t *testing.T) {
	root := newRoot()
	cluster := newCluster(root)
	leaf := newLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixturePermittedHost},
	)
	assert.Nil(t, verify(root, cluster, leaf.Certificate))
}
