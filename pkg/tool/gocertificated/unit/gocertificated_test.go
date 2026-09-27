package unit

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/keystore"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/authority_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
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

func TestCertificateSurvivesArmorRoundTrip(t *testing.T) {
	c := authority_tester.NewRoot().Material().Certificate
	assert.String(
		t,
		c.Subject.CommonName,
		armor.DecodeCertificate(armor.MarshalCertificate(c)).Subject.CommonName,
	)
}

func TestKeySignsAfterArmorRoundTrip(t *testing.T) {
	root := authority_tester.NewRoot()
	directory := filepath.Join(t.TempDir(), "material")
	keystore.Write(directory, root.Material())
	restored := authority.New(keystore.Read(directory))
	cluster := authority_tester.NewCluster(restored)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, authority_tester.Verify(restored, cluster, leaf.Certificate))
}

func TestKeystoreRoundTripPreservesSubject(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "material")
	keystore.Write(directory, authority_tester.NewRoot().Material())
	assert.String(
		t,
		"Example Root CA",
		keystore.Read(directory).Certificate.Subject.CommonName,
	)
}

func TestKeyFileIsNotReadableByOthers(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "material")
	keystore.Write(directory, authority_tester.NewRoot().Material())
	i, e := os.Stat(filepath.Join(directory, constant.KeyFile))

	if e != nil {
		panic(e)
	}

	assert.String(t, "-rw-------", i.Mode().String())
}

func TestIntermediateInheritsRootUnderService(t *testing.T) {
	s, _ := store_tester.NewChain(t)
	defer s.Close()
	c := s.MustAuthority(constant.FixtureClusterAuthority).Material().Certificate
	assert.String(t, "Example Root CA", c.Issuer.CommonName)
}

func TestSecondRootConflicts(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	_, e := v.CreateAuthority(server_tester.RootBody())
	assert.True(t, conflict.Is(e))
}

func TestIssuedLeafVerifiesAgainstTheChain(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, key, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.Nil(t, e)
	assert.StringContains(t, "PRIVATE KEY", key)
	root, cluster := store_tester.ChainOf(s)
	assert.Nil(
		t,
		authority_tester.Verify(
			root,
			cluster,
			armor.DecodeCertificate([]byte(r.Certificate)),
		),
	)
}

func TestIssuedLeafKeyIsNotStored(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.Nil(t, e)
	assert.String(t, "", s.MustBySerial(r.Serial).Key)
}

func TestForeignNameStillFailsThroughTheService(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureImpostor,
			[]string{constant.FixtureForeignDomain},
		),
	)
	assert.Nil(t, e)
	root, cluster := store_tester.ChainOf(s)
	assert.Error(
		t,
		authority_tester.Verify(
			root,
			cluster,
			armor.DecodeCertificate([]byte(r.Certificate)),
		),
	)
}

func TestSignedRequestVerifiesAndStoresNoKey(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, e := v.SignRequest(
		&server.SigningRequestBody{
			Authority: constant.FixtureClusterAuthority,
			Kind:      server.LeafKind(constant.KindServer),
			Request: authority_tester.NewSigningRequest(
				constant.FixtureRequestHost,
			),
		},
	)
	assert.Nil(t, e)
	assert.String(t, "", r.Key)
	root, cluster := store_tester.ChainOf(s)
	assert.Nil(
		t,
		authority_tester.Verify(
			root,
			cluster,
			armor.DecodeCertificate([]byte(r.Certificate)),
		),
	)
}

func TestRevocationListCarriesTheRevokedSerial(t *testing.T) {
	s, v := store_tester.NewChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.Nil(t, e)
	assert.Nil(t, s.Revoke(r.Serial, time.Now()))
	b, f := v.RevocationList(constant.FixtureClusterAuthority)
	assert.Nil(t, f)
	block, _ := pem.Decode(b)
	list, g := x509.ParseRevocationList(block.Bytes)
	assert.Nil(t, g)
	assert.Integer(t, 1, len(list.RevokedCertificateEntries))
	assert.String(
		t,
		r.Serial,
		list.RevokedCertificateEntries[0].SerialNumber.Text(constant.SerialBase),
	)
}

func TestAuthorityRoundTripKeepsSigningKey(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	root := authority_tester.NewRoot()
	s.MustCreate(
		*record.New(constant.KindRoot, constant.RootAuthority, root.Material()),
	)
	restored := authority.New(
		s.MustAuthority(constant.RootAuthority).Material(),
	)
	cluster := authority_tester.NewCluster(restored)
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	assert.Nil(t, authority_tester.Verify(restored, cluster, leaf.Certificate))
}

func TestMissingAuthorityIsNotAnError(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	assert.Nil(t, s.MustAuthority(constant.RootAuthority))
}

func TestLeafKeyIsNotRetained(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	cluster := authority_tester.NewCluster(authority_tester.NewRoot())
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	r := record.New(constant.KindServer, "", leaf)
	assert.String(t, "", r.Key)
}

func TestAuthorityKeyIsRetained(t *testing.T) {
	r := record.New(
		constant.KindRoot,
		constant.RootAuthority,
		authority_tester.NewRoot().Material(),
	)
	assert.StringContains(t, "PRIVATE KEY", r.Key)
}

func TestRevokedAuthorityStopsResolving(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	r := record.New(
		constant.KindRoot,
		constant.RootAuthority,
		authority_tester.NewRoot().Material(),
	)
	s.MustCreate(*r)
	errors.PanicOnError(s.Revoke(r.Serial, time.Now()))
	assert.Nil(t, s.MustAuthority(constant.RootAuthority))
}

func TestExpiringFindsOnlyPastTheHorizon(t *testing.T) {
	s := store_tester.NewStore()
	defer s.Close()
	cluster := authority_tester.NewCluster(authority_tester.NewRoot())
	leaf := authority_tester.NewLeaf(
		cluster,
		constant.FixtureCommonName,
		[]string{constant.FixtureHost},
	)
	s.MustCreate(*record.New(constant.KindServer, "", leaf))
	near, e := s.Expiring(time.Now().AddDate(0, 1, 0))
	errors.PanicOnError(e)
	assert.Integer(t, 0, len(near))
	far, f := s.Expiring(time.Now().AddDate(2, 0, 0))
	errors.PanicOnError(f)
	assert.Integer(t, 1, len(far))
}
