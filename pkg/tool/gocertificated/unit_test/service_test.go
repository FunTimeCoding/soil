package unit_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"testing"
	"time"
)

func newService(s *store.Store) *service.Service {
	return service.New(s, nil)
}

func rootBody() *server.AuthorityBody {
	country := constant.FixtureCountry
	province := constant.FixtureProvince
	organization := constant.FixtureOrganization

	return &server.AuthorityBody{
		Name:         constant.RootAuthority,
		Kind:         server.AuthorityKind(constant.KindRoot),
		CommonName:   constant.FixtureRootCommonName,
		Country:      &country,
		Province:     &province,
		Organization: &organization,
	}
}

func clusterBody() *server.AuthorityBody {
	domain := []string{
		constant.FixtureDomain,
		constant.FixtureInternalDomain,
		constant.FixtureLocalDomain,
	}
	address := []string{constant.FixtureAddress}

	return &server.AuthorityBody{
		Name:             constant.FixtureClusterAuthority,
		Kind:             server.AuthorityKind(constant.KindIntermediate),
		CommonName:       constant.FixtureIssuingCommonName,
		PermittedDomain:  &domain,
		PermittedAddress: &address,
	}
}

func leafBody(common string, host []string) *server.CertificateBody {
	return &server.CertificateBody{
		Authority:  constant.FixtureClusterAuthority,
		Kind:       server.LeafKind(constant.KindServer),
		CommonName: common,
		Host:       &host,
	}
}

func newChain(t *testing.T) (*store.Store, *service.Service) {
	s := newStore()
	v := newService(s)
	_, e := v.CreateAuthority(rootBody())
	assert.Nil(t, e)
	_, f := v.CreateAuthority(clusterBody())
	assert.Nil(t, f)

	return s, v
}

func chainOf(s *store.Store) (*authority.Authority, *authority.Authority) {
	return authority.New(s.MustAuthority(constant.RootAuthority).Material()),
		authority.New(
			s.MustAuthority(constant.FixtureClusterAuthority).Material(),
		)
}

func TestIntermediateInheritsRootUnderService(t *testing.T) {
	s, _ := newChain(t)
	defer s.Close()
	c := s.MustAuthority(constant.FixtureClusterAuthority).Material().Certificate
	assert.String(t, "Example Root CA", c.Issuer.CommonName)
}

func TestSecondRootConflicts(t *testing.T) {
	s, v := newChain(t)
	defer s.Close()
	_, e := v.CreateAuthority(rootBody())
	assert.True(t, errors.Is(e, constant.ErrorConflict))
}

func TestIssuedLeafVerifiesAgainstTheChain(t *testing.T) {
	s, v := newChain(t)
	defer s.Close()
	r, key, e := v.IssueCertificate(
		leafBody(constant.FixtureCommonName, []string{constant.FixtureHost}),
	)
	assert.Nil(t, e)
	assert.StringContains(t, "PRIVATE KEY", key)
	root, cluster := chainOf(s)
	assert.Nil(
		t,
		verify(root, cluster, armor.DecodeCertificate([]byte(r.Certificate))),
	)
}

func TestIssuedLeafKeyIsNotStored(t *testing.T) {
	s, v := newChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		leafBody(constant.FixtureCommonName, []string{constant.FixtureHost}),
	)
	assert.Nil(t, e)
	assert.String(t, "", s.MustBySerial(r.Serial).Key)
}

func TestForeignNameStillFailsThroughTheService(t *testing.T) {
	s, v := newChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		leafBody(
			constant.FixtureImpostor,
			[]string{constant.FixtureForeignDomain},
		),
	)
	assert.Nil(t, e)
	root, cluster := chainOf(s)
	assert.Error(
		t,
		verify(root, cluster, armor.DecodeCertificate([]byte(r.Certificate))),
	)
}

func TestSignedRequestVerifiesAndStoresNoKey(t *testing.T) {
	s, v := newChain(t)
	defer s.Close()
	r, e := v.SignRequest(
		&server.SigningRequestBody{
			Authority: constant.FixtureClusterAuthority,
			Kind:      server.LeafKind(constant.KindServer),
			Request:   newSigningRequest(constant.FixtureRequestHost),
		},
	)
	assert.Nil(t, e)
	assert.String(t, "", r.Key)
	root, cluster := chainOf(s)
	assert.Nil(
		t,
		verify(root, cluster, armor.DecodeCertificate([]byte(r.Certificate))),
	)
}

func TestRevocationListCarriesTheRevokedSerial(t *testing.T) {
	s, v := newChain(t)
	defer s.Close()
	r, _, e := v.IssueCertificate(
		leafBody(constant.FixtureCommonName, []string{constant.FixtureHost}),
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

func newSigningRequest(host string) string {
	k, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	errors.PanicOnError(e)
	t := &x509.CertificateRequest{}
	t.Subject = pkix.Name{CommonName: host}
	t.DNSNames = []string{host}
	b, f := x509.CreateCertificateRequest(rand.Reader, t, k)
	errors.PanicOnError(f)

	return string(
		pem.EncodeToMemory(
			&pem.Block{Type: constant.SigningRequestBlock, Bytes: b},
		),
	)
}
